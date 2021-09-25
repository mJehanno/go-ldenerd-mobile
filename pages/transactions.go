package pages

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/mjehanno/go-ldenerd-mobile/service"
	"github.com/mjehanno/go-ldenerd-mobile/widgets"
)

var TransactionBinding binding.DataList = binding.NewUntypedList()

func GetTransactionsPage(w fyne.Window) *fyne.Container {

	c := container.NewCenter(widget.NewButton("Add Transaction", func() {
		getTransactionFormDialog(w)
	}))

	return c
}

func getTransactionFormDialog(w fyne.Window) {
	transactionInput := widget.NewSelect([]string{"Debit", "Credit"}, func(s string) {})
	transactionTypeFormItem := widget.NewFormItem("Type", transactionInput)

	platInput := widgets.NewNumericalEntry()
	goldInput := widgets.NewNumericalEntry()
	electrumInput := widgets.NewNumericalEntry()
	silverInput := widgets.NewNumericalEntry()
	copperInput := widgets.NewNumericalEntry()
	platFormItem := widget.NewFormItem("Platinum Coins", platInput)
	goldFormItem := widget.NewFormItem("Gold Coins", goldInput)
	electrumFormItem := widget.NewFormItem("Electrum Coins", electrumInput)
	silverFormItem := widget.NewFormItem("Silver Coins", silverInput)
	copperFormItem := widget.NewFormItem("Copper Coins", copperInput)

	reasonInput := widget.NewEntry()
	reasonFormItem := widget.NewFormItem("Reason", reasonInput)

	d := dialog.NewForm("Add a transaction", "Add", "Cancel",
		[]*widget.FormItem{transactionTypeFormItem, copperFormItem, silverFormItem, electrumFormItem, goldFormItem, platFormItem, reasonFormItem}, func(b bool) {

			if b {
				pValue, _ := strconv.Atoi(platInput.Text)
				p := service.Coin{
					Value:    pValue,
					Currency: service.Platinum,
				}
				gValue, _ := strconv.Atoi(goldInput.Text)
				g := service.Coin{
					Value:    gValue,
					Currency: service.Gold,
				}
				eValue, _ := strconv.Atoi(electrumInput.Text)
				e := service.Coin{
					Value:    eValue,
					Currency: service.Electrum,
				}
				sValue, _ := strconv.Atoi(silverInput.Text)
				s := service.Coin{
					Value:    sValue,
					Currency: service.Silver,
				}
				cValue, _ := strconv.Atoi(copperInput.Text)
				c := service.Coin{
					Value:    cValue,
					Currency: service.Copper,
				}
				fmt.Println("transaction type : ", transactionInput.Selected)
				t := service.Transaction{
					Type: service.StringToTransactionType(transactionInput.Selected),
					Amount: []service.Coin{
						p, g, e, s, c,
					},
					Reason: reasonInput.Text,
				}

				service.AddTransaction(t)
				GoldBinding.Set(service.GetGold())
				GoldDetailBinding.Set(service.GetGoldDetail())
			}
		}, w)
	d.Show()
}
