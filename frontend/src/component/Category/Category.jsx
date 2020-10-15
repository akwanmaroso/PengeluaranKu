import React, {useEffect, useState} from "react";
import axios from "axios";
import {isAuth} from "../../utils/isAuth";
import {notifError, notifSuccess} from "../../utils/notifications";


const initialState = {
    name: "",
    description: "",
    color: "",
}
const Category = () => {
    const [state, setState] = useState(initialState)
    const [category, setCategory] = useState([])

    useEffect(() => {
        getCategory()
    }, [])

    const getCategory = () => {
        if (isAuth()) {
            let token = localStorage.getItem("token-user");
            const AuthStr = "Bearer ".concat(token);
            axios
                .get("http://localhost:9000/categories", {
                    headers: {Authorization: AuthStr},
                })
                .then((res) => setCategory(res.data))
                .catch((err) => console.log(err));
        }
    }

    const handleChangeColor = (e) => {
        setState({
            ...state,
            color: e.target.value
        })
    }

    const handleChangeName = (e) => {
        setState({
            ...state,
            name: e.target.value
        })
    }

    const handleChangeDescription = (e) => {
        setState({
            ...state,
            description: e.target.value
        })
    }

    const handleSubmit = (e) => {
        e.preventDefault()
        let token = localStorage.getItem("token-user");
        const AuthStr = "Bearer ".concat(token);
        axios.post("http://localhost:9000/categories", state, {
            headers: {Authorization: AuthStr},
        }).then(() => {
            notifSuccess("Success create category")
            getCategory()
        }).catch(err => notifError(err.message))
    }

    return (
        <div className="container">
            <div className="row">
                <div className="col-4">
                    <h5>Add Category</h5>
                    <hr/>
                    <div className="form-group">
                        <label htmlFor="name">Name</label>
                        <input type="text" className="form-control" id="name" onChange={handleChangeName}/>
                    </div>
                    <div className="form-group">
                        <label htmlFor="description">Description</label>
                        <input type="text" className="form-control" id="description"
                               onChange={handleChangeDescription}/>
                    </div>
                    <div className="form-group">
                        <label htmlFor="color">Color</label>
                        <input type="color" className="form-control" id="color" onChange={handleChangeColor}/>
                    </div>
                    <div className="form-group">
                        <button className="btn btn-success" onClick={handleSubmit}>Create</button>
                    </div>
                </div>
                <div className="col-8">
                    <h5>List Category</h5>
                    <hr/>
                    <table className="table table-striped">
                        <thead className="thead-dark">
                        <tr className="text-center">
                            <th>Color</th>
                            <th>Name</th>
                            <th>Description</th>
                            <th>Action</th>
                        </tr>
                        </thead>
                        <tbody>
                        {
                            category.map((c) => {
                                return (
                                    <tr key={c.id}>
                                        <td className="text-center">
                                            <i style={{color: `${c.color}`}} className="fa fa-circle fa-2x"/>
                                        </td>
                                        <td>{c.name}</td>
                                        <td>{c.description}</td>
                                        <td className="text-center">
                                            <button className="btn btn-danger btn-sm">
                                                <i className="fa fa-remove" />
                                            </button>
                                        </td>
                                    </tr>
                                )
                            })
                        }
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    )
}

export default Category;