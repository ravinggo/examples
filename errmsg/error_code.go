// Code generated by gen-proto-error. DO NOT EDIT.
// source: ./login,./game
package errmsg

import (
	berror "github.com/ravinggo/game/common/berror"
)


func NewBackpackErrorItemNotEnough(extra ...string) *berror.ErrMsg {
	e := &berror.ErrMsg{}
	e.ErrCode = 1
	e.ErrMsg = "backpack_error_item_not_enough"
	if berror.IsOpenStack() {
		e.WithStackTrace()
	}
	if len(extra) > 0 {
		e.ErrExtraInfo = extra[0]
	}
	return e
}

func NewEchoErrorTest1(extra ...string) *berror.ErrMsg {
	e := &berror.ErrMsg{}
	e.ErrCode = 1
	e.ErrMsg = "echo_error_test1"
	if berror.IsOpenStack() {
		e.WithStackTrace()
	}
	if len(extra) > 0 {
		e.ErrExtraInfo = extra[0]
	}
	return e
}

func NewLoginErrorAccountIsNotRegistered(extra ...string) *berror.ErrMsg {
	e := &berror.ErrMsg{}
	e.ErrCode = 1
	e.ErrMsg = "login_error_account_is_not_registered"
	if berror.IsOpenStack() {
		e.WithStackTrace()
	}
	if len(extra) > 0 {
		e.ErrExtraInfo = extra[0]
	}
	return e
}

func NewLoginErrorAccountIsRegistered(extra ...string) *berror.ErrMsg {
	e := &berror.ErrMsg{}
	e.ErrCode = 1
	e.ErrMsg = "login_error_account_is_registered"
	if berror.IsOpenStack() {
		e.WithStackTrace()
	}
	if len(extra) > 0 {
		e.ErrExtraInfo = extra[0]
	}
	return e
}

func NewLoginErrorAccountIsWrong(extra ...string) *berror.ErrMsg {
	e := &berror.ErrMsg{}
	e.ErrCode = 1
	e.ErrMsg = "login_error_account_is_wrong"
	if berror.IsOpenStack() {
		e.WithStackTrace()
	}
	if len(extra) > 0 {
		e.ErrExtraInfo = extra[0]
	}
	return e
}

func NewLoginErrorAgeLimit(extra ...string) *berror.ErrMsg {
	e := &berror.ErrMsg{}
	e.ErrCode = 1
	e.ErrMsg = "login_error_age_limit"
	if berror.IsOpenStack() {
		e.WithStackTrace()
	}
	if len(extra) > 0 {
		e.ErrExtraInfo = extra[0]
	}
	return e
}

func NewLoginErrorAuthenticationNotApproved(extra ...string) *berror.ErrMsg {
	e := &berror.ErrMsg{}
	e.ErrCode = 1
	e.ErrMsg = "login_error_authentication_not_approved"
	if berror.IsOpenStack() {
		e.WithStackTrace()
	}
	if len(extra) > 0 {
		e.ErrExtraInfo = extra[0]
	}
	return e
}

func NewLoginErrorFeedbackContactTooLong(extra ...string) *berror.ErrMsg {
	e := &berror.ErrMsg{}
	e.ErrCode = 1
	e.ErrMsg = "login_error_feedback_contact_too_long"
	if berror.IsOpenStack() {
		e.WithStackTrace()
	}
	if len(extra) > 0 {
		e.ErrExtraInfo = extra[0]
	}
	return e
}

func NewLoginErrorFeedbackContentTooLong(extra ...string) *berror.ErrMsg {
	e := &berror.ErrMsg{}
	e.ErrCode = 1
	e.ErrMsg = "login_error_feedback_content_too_long"
	if berror.IsOpenStack() {
		e.WithStackTrace()
	}
	if len(extra) > 0 {
		e.ErrExtraInfo = extra[0]
	}
	return e
}

func NewLoginErrorFormatError(extra ...string) *berror.ErrMsg {
	e := &berror.ErrMsg{}
	e.ErrCode = 1
	e.ErrMsg = "login_error_format_error"
	if berror.IsOpenStack() {
		e.WithStackTrace()
	}
	if len(extra) > 0 {
		e.ErrExtraInfo = extra[0]
	}
	return e
}

func NewLoginErrorInvalidLoginType(extra ...string) *berror.ErrMsg {
	e := &berror.ErrMsg{}
	e.ErrCode = 1
	e.ErrMsg = "login_error_invalid_login_type"
	if berror.IsOpenStack() {
		e.WithStackTrace()
	}
	if len(extra) > 0 {
		e.ErrExtraInfo = extra[0]
	}
	return e
}

func NewLoginErrorLimitPlayTimeConfigWrong(extra ...string) *berror.ErrMsg {
	e := &berror.ErrMsg{}
	e.ErrCode = 1
	e.ErrMsg = "login_error_limit_play_time_config_wrong"
	if berror.IsOpenStack() {
		e.WithStackTrace()
	}
	if len(extra) > 0 {
		e.ErrExtraInfo = extra[0]
	}
	return e
}

func NewLoginErrorNotRealNameAuthentication(extra ...string) *berror.ErrMsg {
	e := &berror.ErrMsg{}
	e.ErrCode = 1
	e.ErrMsg = "login_error_not_real_name_authentication"
	if berror.IsOpenStack() {
		e.WithStackTrace()
	}
	if len(extra) > 0 {
		e.ErrExtraInfo = extra[0]
	}
	return e
}

func NewLoginErrorPasswordFormatError(extra ...string) *berror.ErrMsg {
	e := &berror.ErrMsg{}
	e.ErrCode = 1
	e.ErrMsg = "login_error_password_format_error"
	if berror.IsOpenStack() {
		e.WithStackTrace()
	}
	if len(extra) > 0 {
		e.ErrExtraInfo = extra[0]
	}
	return e
}

func NewLoginErrorTokenExpired(extra ...string) *berror.ErrMsg {
	e := &berror.ErrMsg{}
	e.ErrCode = 1
	e.ErrMsg = "login_error_token_expired"
	if berror.IsOpenStack() {
		e.WithStackTrace()
	}
	if len(extra) > 0 {
		e.ErrExtraInfo = extra[0]
	}
	return e
}

func NewLoginErrorTokenExpiredIsReconnect(extra ...string) *berror.ErrMsg {
	e := &berror.ErrMsg{}
	e.ErrCode = 1
	e.ErrMsg = "login_error_token_expired_is_reconnect"
	if berror.IsOpenStack() {
		e.WithStackTrace()
	}
	if len(extra) > 0 {
		e.ErrExtraInfo = extra[0]
	}
	return e
}

