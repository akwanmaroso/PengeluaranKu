import React, {useState} from "react";
import axios from "axios";
import {notifError, notifSuccess} from "../../utils/notifications";

const initialData = {
    amount: 0,
    description: "",
    in_out: false,
    category_id: 1,
    partner_id: 1,
    creator_id: 1,
};
const AddTransactions = (props) => {
    const [data, setData] = useState(initialData);

    const handleChangeAmount = (e) => {
        setData({
            ...data,
            amount: Number(e.target.value),
        });
    };

    const handleChangeDesc = (e) => {
        setData({
            ...data,
            description: e.target.value,
        });
    };

    const postTransaction = (token) => {
        axios
            .post("http://localhost:9000/transactions", data, {
                headers: {Authorization: token},
            })
            .then(() => {
                notifSuccess("Add income success");
                setData(initialData);
            })
            .catch((err) => {
                notifError(`Failed to add income ${err.message}!`);
            });
        props.getData()
    }

    const handleAddIncome = (e) => {
        e.preventDefault();


        console.log(data)
        let token = localStorage.getItem("token-user");
        const AuthStr = "Bearer ".concat(token);
        console.log(data);
        postTransaction(AuthStr);
        resetForm();
    };

    const handleAddSpending = (e) => {
        e.preventDefault();
        data.in_out = true
        console.log(data)

        let token = localStorage.getItem("token-user");
        const AuthStr = "Bearer ".concat(token);
        postTransaction(AuthStr)
        resetForm();
    };

    const resetForm = () => {
        document.getElementById("form-transactions").reset();
    };

    const validateForm = () => {
        return data.description === "" || data.amount === 0;
    };

    return (
        <div>
            <form id="form-transactions">
                <input
                    name="amount"
                    type="number"
                    className="form-control mb-2"
                    placeholder="eg. 800000"
                    onChange={handleChangeAmount}
                    required
                />
                <input
                    name="description"
                    type="text"
                    className="form-control"
                    placeholder="eg. Buy whiskies"
                    onChange={handleChangeDesc}
                    required
                />
                <div className="row mt-3">
                    <div className="col">
                        <button
                            type="submit"
                            disabled={validateForm()}
                            onClick={handleAddIncome}
                            className="btn btn-success float-right"
                        >
                            Add Income
                        </button>
                        <button
                            type="submit"
                            disabled={validateForm()}
                            onClick={handleAddSpending}
                            className="btn btn-danger float-right mr-2"
                        >
                            Add Spending
                        </button>
                    </div>
                </div>
            </form>
        </div>
    );
};

export default AddTransactions;
