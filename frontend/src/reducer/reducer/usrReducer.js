const initState = {
    user: null,
    isLoading: false,
    error: null,
}

const usrReducer = (state = initState, action) => {
    switch (action.type) {
        case "Loading_User":
            return { ...state, isLoading: true }
        case "Register_User":
            return { ...state, user: action.payload, isLoading: false }
        case "Login_User":
            return { ...state, user: action.payload, isLoading: false }
        case "LOGOUT":
            localStorage.removeItem("access_token")
            return { ...state, user: null }
        case "ERROR":
            return { ...state, error: action.payload, isLoading: false }
        default:
            return state
    }
}

export default usrReducer
