import React, { useEffect, useState } from "react";
import axios from "axios";
import { notifError, notifSuccess } from "../../utils/notifications";
import { isAuth } from "../../utils/isAuth";

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
  const [category, setCategory] = useState([]);

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

  const handleChangeCategory = (e) => {
    setData({
      ...data,
      category_id: Number(e.target.value),
    });
  };

  const postTransaction = (token, type) => {
    axios
      .post("http://localhost:9000/transactions", data, {
        headers: { Authorization: token },
      })
      .then(() => {
        notifSuccess(`Add ${type} success`);
        setData(initialData);
      })
      .catch((err) => {
        notifError(`Failed to add income ${err.message}!`);
      });
    props.getData();
  };

  const handleAddIncome = (e) => {
    e.preventDefault();

    let token = localStorage.getItem("token-user");
    const AuthStr = "Bearer ".concat(token);
    postTransaction(AuthStr, "income");
    resetForm();
  };

  const handleAddSpending = (e) => {
    e.preventDefault();
    data.in_out = true;

    let token = localStorage.getItem("token-user");
    const AuthStr = "Bearer ".concat(token);
    postTransaction(AuthStr, "spending");
    resetForm();
  };

  const getCategory = () => {
    if (isAuth()) {
      let token = localStorage.getItem("token-user");
      const AuthStr = "Bearer ".concat(token);
      axios
        .get("http://localhost:9000/categories", {
          headers: { Authorization: AuthStr },
        })
        .then((res) => setCategory(res.data))
        .catch((err) => console.log(err));
    }
  };

  useEffect(() => {
    getCategory();
  }, []);

  const resetForm = () => {
    document.getElementById("form-transactions").reset();
  };

  const validateForm = () => {
    return data.description === "" || data.amount === 0;
  };

  return (
    <div>
      <form id="form-transactions">
        <div className="form-group">
          <input
            name="amount"
            type="number"
            className="form-control mb-2"
            placeholder="eg. 800000"
            onChange={handleChangeAmount}
            required
          />
        </div>
        <div className="form-group">
          <input
            name="description"
            type="text"
            className="form-control"
            placeholder="eg. Buy whiskies"
            onChange={handleChangeDesc}
            required
          />
        </div>
        <div className="form-group">
          <select
            name="category"
            id="category"
            className="form-control"
            onChange={handleChangeCategory}
          >
            {/*<option value="as">Select Category</option>*/}
            {category.map((c) => {
              return (
                <option key={c.id} value={c.id}>
                  {c.name}
                </option>
              );
            })}
          </select>
        </div>
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
