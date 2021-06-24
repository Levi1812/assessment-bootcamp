import api from "../../API/api"

const access_token = !localStorage.getItem("access_token") ? "" : localStorage.getItem("access_token")

export const fetchPass = () => {
    return async (dispatch) => {
        try {
            dispatch({ type: "Password_Loading" })

            const response = await api({
                method: "GET",
                url: "/site",
                headers: {
                    "authorization": access_token
                }
            })

            console.log(response.data)

            return dispatch({ type: "Password_Loading", payload: response.data })

        } catch (err) {
            dispatch({ type: "ERROR" })
            console.log(err.response)
        }
    }
}


export const createPass = (payload) => {
    return async (dispatch) => {
        try {
            dispatch({ type: "Password_Loading" })

            const response = await api({
                method: "POST",
                url: "/site",
                data: payload,
                headers: {
                    "authorization": access_token
                }
            })

            console.log(response.data)

            return dispatch({ type: "Create_Password", payload: response.data })

        } catch (err) {
            dispatch({ type: "ERROR" })
            console.log(err.response)
        }
    }

}

export const getSite = (id, payload) => {
    return async (dispatch) => {
        try {
            dispatch({ type: "Password_Loading" })

            const response = await api({
                method: "GET",
                url: `/site/${id}`,
                data: payload,
                headers: {
                    "authorization": access_token
                }
            })

            console.log(response.data)

            return dispatch({ type: "Get_SiteID", payload: response.data })

        } catch (err) {
            dispatch({ type: "ERROR" })
            console.log(err.response)
        }
    }
}

export const updatePass = (id, payload) => {
    return async (dispatch) => {
        try {
            dispatch({ type: "Password_Loading" })

            const response = await api({
                method: "PUT",
                url: `/site/${id}`,
                data: payload,
                headers: {
                    "authorization": access_token
                }
            })

            console.log(response.data)

            return dispatch({ type: "Update_Password", payload: response.data })

        } catch (err) {
            dispatch({ type: "ERROR" })
            console.log(err.response)
        }
    }
}


export const deletePass = (id, history) => {
    return async (dispatch) => {
        try {
            dispatch({ type: "Password_Loading" })

            const response = await api({
                method: "DELETE",
                url: `/site/${id}`,
                headers: {
                    "authorization": access_token
                }
            })
            history.push("/pass")

            console.log(response.data)

            return dispatch({ type: "Delete_Password", payload: response.data })
        } catch (err) {
            dispatch({ type: "ERROR" })
            console.log(err.response)
        }
    }
}
