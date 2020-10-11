export function isAuth() {
    if (localStorage.getItem("token-user") === null) {
        return false
    }
    return true
}