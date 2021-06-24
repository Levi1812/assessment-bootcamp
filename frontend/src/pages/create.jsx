import React from 'react'

import Navbar from "../components/navigationbar"
import { useDispatch, useSelector } from 'react-redux'
import { useState } from 'react'

import { useHistory } from "react-router-dom"
import { createPass, fetchPass } from "../reducer/act/pwAction"

function Create() {
    const history = useHistory()
    const dispatch = useDispatch()

    const [website, setWebsite] = useState("")
    const [pass, setPass] = useState("")

    const { error } = useSelector(state => state.user)

    const submitPass = (e) => {
        e.preventDefault()
        const data = {
            Website: website,
            Pass: pass,
        }

        console.log(data)


        if (!error) {
            dispatch(createPass(data))
            dispatch(fetchPass)
            history.push("/pass")
        }

    }


    return (
        <div>
            <Navbar />
            <div className="container">
                <form style={{ textAlign: "center", paddingTop: "150px", paddingBottom: "150px" }} onSubmit={submitPass}>
                    <h2>Create New Password</h2>
                    <div className="mb-3">
                        <label for="title" className="form-label">Website</label>
                        <input type="text" className="form-control" id="title" onChange={e => {
                            e.preventDefault()
                            setWebsite(e.target.value)
                        }} />
                    </div>
                    <div className="mb-3">
                        <label for="auhtor" className="form-label">Password</label>
                        <input type="text" className="form-control" id="auhtor" onChange={e => {
                            e.preventDefault()
                            setPass(e.target.value)
                        }} />
                    </div>
                    <button type="submit" className="btn btn-primary" style={{ margin: "0px 100px" }}>Create</button>
                </form>
            </div>
            {/* <Footer/> */}
        </div>
    )
}


export default Create