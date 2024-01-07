package application

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/repository"
	"github.com/Furkan-Gulsen/Checkout-System/internal/interfaces/dto"
	"github.com/Furkan-Gulsen/Checkout-System/pkg/utils"
)

var _ CartAppInterface = &cartApp{}

type cartApp struct {
	cartRepo     repository.CartRepositoryI
	itemApp      ItemAppInterface
	vasItemApp   VasItemAppInterface
	promotionApp PromotionAppInterface
}

func NewCartApp(cartRepo repository.CartRepositoryI, itemApp ItemAppInterface, vasItemApp VasItemAppInterface, promotionApp PromotionAppInterface) *cartApp {
	return &cartApp{
		cartRepo:     cartRepo,
		itemApp:      itemApp,
		vasItemApp:   vasItemApp,
		promotionApp: promotionApp,
	}
}

type CartAppInterface interface {
	ApplyPromotion(cartId int, promotionId int) (*entity.Cart, error) // OK
	Display(cartId int) (*dto.DisplayCartDTO, error)
	ResetCart(cartId int) error
	AddItem(cartId int, item *entity.Item) (*entity.Cart, error)        // OK
	UpdateCartPriceAndQuantity(cart *entity.Cart) (*entity.Cart, error) // OK
}

func (app *cartApp) ApplyPromotion(cartId int, promotionId int) (*entity.Cart, error) {
	cart, err := app.cartRepo.GetByID(cartId)
	if err != nil {
		slog.Error("Cart not found. Error: ", err)
		return nil, fmt.Errorf("cart not found. CartID: %d", cartId)
	}

	if cart.AppliedPromotionId != 0 {
		return nil, fmt.Errorf("promotion already applied")
	}

	promotion, err := app.promotionApp.GetById(promotionId)
	if err != nil {
		slog.Error("Promotion not found. Error: ", err)
		return nil, fmt.Errorf("promotion is not found, promotionId: %d", promotionId)
	}

	items, _ := app.itemApp.ListByCartId(cartId)
	if len(items) > 0 {
		cart = calcCartPricesWithPromotion(cart, items, promotion)
		app.cartRepo.Update(cart)
	}

	cart.AppliedPromotionId = promotion.Id
	_, err = app.cartRepo.Update(cart)
	if err != nil {
		slog.Error("Failed to apply promotion. Error: ", err)
		return nil, fmt.Errorf("failed to apply promotion, cartId: %d, promotionId: %d", cartId, promotionId)
	}

	return cart, nil
}

func calcCartPricesWithPromotion(cart *entity.Cart, items []*entity.Item, promotion *entity.Promotion) *entity.Cart {
	fmt.Println("-------------------------------------")
	totalDiscount := float64(0)
	firstSellerID := items[0].SellerID
	sameSeller := true
	sameSellerTotalDiscount := float64(0)

	for _, item := range items {
		if promotion.PromotionType == entity.CategoryPromotion && item.CategoryID == promotion.CategoryP.CategoryID && item.ItemType == entity.DefaultItem {
			itemDiscount := item.Price * (float64(promotion.CategoryP.DiscountRate) / 100)
			totalDiscount += float64(item.Quantity) * itemDiscount
		} else if promotion.PromotionType == entity.SameSellerPromotion {
			if item.SellerID != firstSellerID {
				sameSeller = false
				break
			}

			itemDiscount := item.Price * (float64(promotion.SameSellerP.DiscountRate) / 100)
			sameSellerTotalDiscount += float64(item.Quantity) * itemDiscount
		}
	}

	if promotion.PromotionType == entity.TotalPricePromotion {
		for _, rnge := range promotion.TotalPriceP {
			fmt.Println("cart.TotalAmount: ", cart.TotalAmount)
			if cart.TotalPrice >= rnge.PriceRangeStart && cart.TotalPrice <= rnge.PriceRangeEnd {
				totalDiscount = rnge.DiscountAmount
				fmt.Println("totalDiscount: ", totalDiscount)
				break
			}
		}
	} else if promotion.PromotionType == entity.SameSellerPromotion && sameSeller {
		totalDiscount = sameSellerTotalDiscount
	}

	fmt.Println("* totalDiscount: ", totalDiscount)
	cart.TotalDiscount = totalDiscount
	cart.TotalAmount = cart.TotalPrice - totalDiscount

	fmt.Println("-------------------------------------")
	return cart
}

func (app *cartApp) Display(cartId int) (*dto.DisplayCartDTO, error) {
	cart, err := app.cartRepo.GetByID(cartId)
	if err != nil {
		return nil, fmt.Errorf("cart not found. CartID: %d", cartId)
	}

	items, itemErr := app.itemApp.ListByCartId(cartId)
	if itemErr != nil {
		return nil, fmt.Errorf("failed to retrieve cart items. Error: %v", itemErr)
	}

	var itemDTOs []*dto.ItemDTO
	for _, item := range items {
		vasItems, vasItemErr := app.vasItemApp.ListByItemId(item.Id)
		if vasItemErr != nil {
			return nil, fmt.Errorf("failed to retrieve vas items. Error: %v", vasItemErr)
		}

		itemDTOs = append(itemDTOs, &dto.ItemDTO{
			ID:         item.Id,
			CategoryID: item.CategoryID,
			SellerID:   item.SellerID,
			CartID:     item.CartID,
			Price:      item.Price,
			ItemType:   item.ItemType,
			VasItems:   vasItems,
		})
	}

	displayCartDTO := &dto.DisplayCartDTO{
		ID:                 cart.Id,
		TotalPrice:         cart.TotalPrice,
		TotalDiscount:      cart.TotalDiscount,
		TotalAmount:        cart.TotalAmount,
		AppliedPromotionId: cart.AppliedPromotionId,
		Items:              itemDTOs,
	}

	return displayCartDTO, nil
}

func (app *cartApp) ResetCart(cartId int) error {
	err := app.cartRepo.Delete(cartId)
	if err != nil {
		return fmt.Errorf("failed to reset cart: %v", err)
	}

	items, err := app.itemApp.ListByCartId(cartId)
	if err != nil {
		return fmt.Errorf("failed to reset cart: %v", err)
	}

	var wg sync.WaitGroup
	for _, item := range items {
		wg.Add(1)
		go func(item *entity.Item) {
			defer wg.Done()

			app.itemApp.Delete(item.Id)
			vasItems, _ := app.vasItemApp.ListByItemId(item.Id)
			for _, vasItem := range vasItems {
				app.vasItemApp.DeleteById(vasItem.Id)
			}
		}(item)
	}
	wg.Wait()

	return nil
}

func (app *cartApp) AddItem(cartId int, item *entity.Item) (*entity.Cart, error) {
	cart, err := app.cartRepo.GetByID(cartId)
	if cart == nil || err != nil {
		cart = &entity.Cart{
			Id: utils.GenerateID(),
		}
		_, err := app.cartRepo.Create(cart)
		if err != nil {
			return nil, fmt.Errorf("failed create cart error: %v", err)
		}
	}
	item.CartID = cartId

	itemValidateErr := item.Validate()
	if itemValidateErr != nil {
		return nil, fmt.Errorf("failed to add item: %v", itemValidateErr)
	}

	item, err = app.itemApp.Create(item)
	if err != nil {
		return nil, fmt.Errorf("failed to add item: %v", err)
	}

	cart, updCartErr := app.UpdateCartPriceAndQuantity(cart)
	if updCartErr != nil {
		app.itemApp.Delete(item.Id) // * Rollback
		slog.Error("Failed to update cart price and quantity. Error: ", updCartErr)
		return nil, fmt.Errorf("failed to update cart price and quantity. Error: %v", updCartErr)
	}

	return cart, nil
}

func (app *cartApp) UpdateCartPriceAndQuantity(cart *entity.Cart) (*entity.Cart, error) {
	items, listItemsErr := app.itemApp.ListByCartId(cart.Id)
	if listItemsErr != nil {
		return nil, fmt.Errorf("list items error: %v", listItemsErr)
	}

	var totalPrice float64
	var totalQuantity int

	for _, item := range items {
		// fmt.Println("ID: ", item.Id)
		totalQuantity += item.Quantity
		totalPrice += item.Price * float64(item.Quantity)
		// fmt.Println("totalPrice: ", totalPrice)
		// fmt.Println("totalQuantity: ", totalQuantity)
		vasItems, listVasItemErr := app.vasItemApp.ListByItemId(item.Id)
		if len(vasItems) > 0 && listVasItemErr == nil {
			for _, vasItem := range vasItems {
				totalPrice += vasItem.Price * float64(vasItem.Quantity)
				// fmt.Println("--------------------")
				// fmt.Println("vasItem_price: ", vasItem.Price)
				// fmt.Println("vasItem_quantity: ", vasItem.Quantity)
				// fmt.Println("vasItem_totalPrice: ", vasItem.Price*float64(vasItem.Quantity))
				// fmt.Println("totalPrice: ", totalPrice)
				// fmt.Println("--------------------")
			}
		}

		// fmt.Println("..............................")
	}

	cart.TotalPrice = totalPrice
	fmt.Println("2. totalPrice: ", totalPrice)
	fmt.Println("2. totalQuantity: ", totalQuantity)

	if totalQuantity > 30 {
		return nil, fmt.Errorf("total quantity can not be more than 30. Total Quantity: %d", totalQuantity)
	}

	if cart.AppliedPromotionId != 0 {
		promotion, getPromErr := app.promotionApp.GetById(cart.AppliedPromotionId)
		if getPromErr != nil {
			return nil, fmt.Errorf("get promotion error: %v", getPromErr)
		}
		cart = calcCartPricesWithPromotion(cart, items, promotion)
		fmt.Println("1 TotalAmount: ", cart.TotalAmount)
		fmt.Println("1 TotalDiscount: ", cart.TotalDiscount)
		fmt.Println("1 TotalPrice: ", cart.TotalPrice)
		fmt.Println("1 AppliedPromotionId: ", cart.AppliedPromotionId)
	} else {
		cart.TotalAmount = totalPrice
		cart.TotalDiscount = 0
	}

	if cart.TotalAmount > 500000 {
		return nil, fmt.Errorf("total amount can not be more than 500000. Total Amount: %f", cart.TotalAmount)
	}

	_, err := app.cartRepo.Update(cart)
	if err != nil {
		return nil, fmt.Errorf("update cart error: %v", err)
	}

	return cart, nil
}
