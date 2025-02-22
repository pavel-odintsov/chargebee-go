package enum

type PaymentMethod string

const (
	PaymentMethodCard                  PaymentMethod = "card"
	PaymentMethodCash                  PaymentMethod = "cash"
	PaymentMethodCheck                 PaymentMethod = "check"
	PaymentMethodChargeback            PaymentMethod = "chargeback"
	PaymentMethodBankTransfer          PaymentMethod = "bank_transfer"
	PaymentMethodAmazonPayments        PaymentMethod = "amazon_payments"
	PaymentMethodPaypalExpressCheckout PaymentMethod = "paypal_express_checkout"
	PaymentMethodDirectDebit           PaymentMethod = "direct_debit"
	PaymentMethodAlipay                PaymentMethod = "alipay"
	PaymentMethodUnionpay              PaymentMethod = "unionpay"
	PaymentMethodApplePay              PaymentMethod = "apple_pay"
	PaymentMethodWechatPay             PaymentMethod = "wechat_pay"
	PaymentMethodAchCredit             PaymentMethod = "ach_credit"
	PaymentMethodSepaCredit            PaymentMethod = "sepa_credit"
	PaymentMethodIdeal                 PaymentMethod = "ideal"
	PaymentMethodGooglePay             PaymentMethod = "google_pay"
	PaymentMethodSofort                PaymentMethod = "sofort"
	PaymentMethodBancontact            PaymentMethod = "bancontact"
	PaymentMethodGiropay               PaymentMethod = "giropay"
	PaymentMethodDotpay                PaymentMethod = "dotpay"
	PaymentMethodOther                 PaymentMethod = "other"
	PaymentMethodUpi                   PaymentMethod = "upi"
	PaymentMethodNetbankingEmandates   PaymentMethod = "netbanking_emandates"
	PaymentMethodCustom                PaymentMethod = "custom"
	PaymentMethodBoleto                PaymentMethod = "boleto"
)
