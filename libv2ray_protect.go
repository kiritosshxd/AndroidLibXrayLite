package libv2ray

import (
    "log"

    v2internet "github.com/xtls/xray-core/transport/internet"
)

func attachProtectedDialer(cb CoreCallbackHandler) {
    if cb == nil {
        return
    }

    d := NewProtectedDialer(cb) // vem do libv2ray_support.go que você copiou
    d.PrepareResolveChan()

    v2internet.UseAlternativeSystemDialer(d)

    log.Println("[xray-lite] protected dialer enabled")
}
