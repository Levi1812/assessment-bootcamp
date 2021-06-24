import { deletePass, fetchPass } from "../reducer/act/pwAction"
import Navbar from "../components/navigationbar"
import { useDispatch, useSelector } from "react-redux"
import React, { useEffect } from "react"
import { useHistory } from "react-router-dom"


function Password() {
    const history = useHistory()
    const dispatch = useDispatch()

    const { passw } = useSelector(state => state.pass)


    useEffect(() => {
        dispatch(fetchPass())
    }, [])

    console.log(passw.data)

    const create = () => {
        history.push("/createPass")
    }

    return (
        <div>
            <Navbar />
            <div style={{ padding: "10%" }}>
                <button class="btn btn-primary" onClick={create}>Create New Book</button>
                <table class="table table-hover">
                    <thead>
                        <tr>
                            <th>NO</th>
                            <th scope="col">Website</th>
                            <th scope="col">Password</th>
                            <th scope="col">Update</th>
                            <th scope="col">Delete</th>
                        </tr>
                    </thead>
                    <tbody>
                        {passw && passw.map((pass, index) => {
                            return (
                                <tr>
                                    <td>{index + 1}</td>
                                    <td>{pass.website}</td>
                                    <td>{pass.password}</td>
                                    <td>
                                        <button class="btn btn-warning" onClick={(e) => {
                                            e.preventDefault()
                                            history.push("/updatePass/" + pass.ID)
                                        }}>Update</button>
                                    </td>
                                    <td>
                                        <button class="btn btn-danger" onClick={(e) => {
                                            e.preventDefault()
                                            dispatch(deletePass(pass.ID, history))
                                            dispatch(fetchPass())
                                        }}>Delete</button>
                                    </td>
                                </tr>
                            )
                        })}
                    </tbody>
                </table>
            </div>
        </div>
    )
    // }
}


export default Password