import React, { useState, useEffect } from "react";
import { useDispatch } from "react-redux";
import { logoutUser } from "../reducer/act/usrAction";
import { useHistory, useLocation } from "react-router-dom";
import { Navbar, Nav } from "react-bootstrap"

const Navigation = () => {
    const [pageURL, setPageURL] = useState("");
    const authUser = !!localStorage.getItem("access_token");
    const location = useLocation();
    const dispatch = useDispatch();
    const history = useHistory();

    useEffect(() => {
        setPageURL(location.pathname);
    }, []);

    const handlerLogout = () => {
        dispatch(logoutUser());
        history.push("/login");
    };

    return (
        <>
            <Navbar collapseOnSelect expand="lg" bg="light" variant="light">
                <div className="container">

                    <Navbar.Brand href="/" className="ml-5">
                        Password Manager
                    </Navbar.Brand>
                    <Navbar.Toggle aria-controls="responsive-navbar-nav" />
                    <Navbar.Collapse
                        id="responsive-navbar-nav"
                        className="justify-content-end"
                    >
                        <Nav>
                            {authUser && (
                                <>
                                    <Nav.Link className="mr-5" onClick={handlerLogout}>
                                        Logout
                                    </Nav.Link>
                                </>
                            )}
                            {!authUser && (
                                <>
                                    {pageURL === "/register" || (
                                        <Nav.Link href="/register" className="mr-5">
                                            Register
                                        </Nav.Link>
                                    )}
                                    {pageURL === "/login" || (
                                        <Nav.Link href="/login" className="mr-5">
                                            Login
                                        </Nav.Link>
                                    )}
                                </>
                            )}
                        </Nav>
                    </Navbar.Collapse>
                </div>
            </Navbar>
        </>
    );
};

export default Navigation;