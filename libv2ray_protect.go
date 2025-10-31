package libv2ray

import (
    "log"

    v2internet "github.com/xtls/xray-core/transport/internet"
)

// Essa função será chamada no NewCoreController
func attachProtectedDialer(cb CoreCallbackHandler) {
    if cb == nil {
        return
    }

    // cb tem Protect(int) bool (a gente vai recolocar na interface)
    d := NewProtectedDialer(cb)
    // essa função existia no código da 25.3.31
    d.PrepareResolveChan()

    // registra o dialer protegido no xray
    v2internet.UseAlternativeSystemDialer(d)

    log.Println("[xray-lite] protected dialer enabled (vpn protect)")
}
