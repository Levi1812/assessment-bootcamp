import api from "../../API/api"

export const register = (payload) => {
    return async (dispatch) => {
        try {
            dispatch({ type: "Loading_User" })

            const response = await api({
                method: "POST",
                url: "/users/register",
                data: payload,
            })

            console.log(response.data)

            return dispatch({ type: "User_Register", payload: response.data })

        } catch (err) {
            console.log(err.response)
        }
    }
}

export const loginusr = (payload, history) => {

    return async (dispatch) => {
        try {
            dispatch({ type: "Loading_User" })

            const response = await api({
                method: "POST",
                url: "/users/login",
                data: payload,
            })

            console.log("inirespon", response.data)

            localStorage.setItem("access_token", response.data.authorization)

            // console.log(payload)

            history.push("/site")

            return dispatch({ type: "Login_User", payload: response.data })


        } catch (err) {
            console.log(payload)
            dispatch({ type: "ERROR", payload: err })
            console.log(err.response)
        }
    }
}

export const logoutUser = () => ({
    type: "LOGOUT"
})
