package auth

import "github.com/eugeis/gee/net"

func (o *AuthEventhorizonInitializer) ActivatePasswordEncryption() {
	o.AccountAggregateInitializer.AddCreatePreparer(
		func(cmd *CreateAccount, entity *Account) (err error) {
			cmd.Password, err = net.Encrypt(cmd.Password)
			return
		})

	o.AccountAggregateInitializer.AddUpdatePreparer(
		func(cmd *UpdateAccount, entity *Account) (err error) {
			if len(cmd.Password) > 0 {
				cmd.Password, err = net.Encrypt(cmd.Password)
			}
			return
		})
}
