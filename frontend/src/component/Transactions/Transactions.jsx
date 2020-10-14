import React, {useEffect, useState} from "react";
import {notifError, notifSuccess} from "../../utils/notifications";
import {convertToRupiah} from "../helpers/helpers"
import {isAuth} from "../../utils/isAuth"
import axios from "axios";
import Summary from "../Summary";
import moment from "moment";
import AddTransactions from "./AddTransactions";

const Transactions = () => {
    const [state, setState] = useState({
        transaction: [],
    });

    const handleDelete = async (id) => {
        try {
            const res = await axios.delete(
                `http://localhost:9000/transactions/${id}`
            );
            console.log(res)
            getData()
            notifSuccess("Delete success");
        } catch (error) {
            notifError(`Delete failed ${error.message}`);
        }
    };

    const getData = () => {
        if (isAuth()) {
            let token = localStorage.getItem("token-user");
            const AuthStr = "Bearer ".concat(token);
            axios
                .get("http://localhost:9000/transactions", {
                    headers: {Authorization: AuthStr},
                })
                .then((res) => setState({transaction: res.data}))
                .catch((err) => console.log(err));
        }
    }

    useEffect(() => {
        getData()
    }, []);


    return (
        <div className="container">
            <div className="row">
                <div className="col">
                    <AddTransactions getData={getData}/>
                </div>
            </div>
            <Summary data={state}/>
            <div className="row">
                <div className="col">
                    <div className="table-responsive">
                        <table className="table">
                            <thead>
                            <tr>
                                <th>No</th>
                                <th>Date</th>
                                <th>Transaction Description</th>
                                <th>Amount</th>
                                <th>Action</th>
                            </tr>
                            </thead>
                            <tbody>
                            {state.transaction.map((tr, index) => {
                                return (
                                    <tr key={tr.id}>
                                        <td>{index + 1}</td>
                                        <td>{moment(tr.date).calendar()}</td>
                                        <td>
                                            {tr.description}{" "}
                                            <span
                                                className="badge float-right badge-pill"
                                                style={{
                                                    backgroundColor: `${tr.category.color}`,
                                                    color: "white",
                                                }}
                                            > {tr.category.name}
                                                          </span>
                                        </td>
                                        <td>
                                            {tr.in_out
                                                ? "-" + convertToRupiah(tr.amount)
                                                : "+" + convertToRupiah(tr.amount)}
                                        </td>
                                        <td>
                                            <button
                                                onClick={() => handleDelete(tr.id)}
                                                className="btn btn-danger btn-sm text-center"
                                            >
                                                Remove
                                            </button>
                                        </td>
                                    </tr>
                                );
                            })}
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Transactions;