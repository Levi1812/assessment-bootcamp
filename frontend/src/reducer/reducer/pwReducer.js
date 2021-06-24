const initState = {
    passw: [],
    pass: null,
    isLoading: false,
    error: null
}

const pwReducer = (state = initState, action) => {
    switch (action.type) {
        case "Loading_Password":
            return { ...state, isLoading: true }
        case "Fetch":
            return { ...state, isLoading: false, passw: action.payload }
        case "Create_Password":
            return { ...state, isLoading: false, pass: action.payload }
        case "Update_Password":
            return { ...state, isLoading: false, pass: action.payload }
        case "Delete_Password":
            return { ...state, isLoading: false, pass: null }
        case "ERROR":
            return { ...state, isLoading: false }
        default:
            return state
    }
}

export default pwReducer