import React from "react";
import {convertToRupiah} from "./helpers/helpers"

const Summary = (props) => {
  // Calculate income
  const income = props.data.transaction.reduce((items, tr) => {
    return tr.in_out === true ? items : items + tr.amount;
  }, 0);

  // Calculate spending
  const spending = props.data.transaction.reduce((items, tr) => {
    return tr.in_out === false ? items : items + tr.amount;
  }, 0);

  // Calculate difference
  const difference = income - spending;

  return (
    <div>
      <div className="row text-center mb-3">
        <div className="col-md-4 mt-4">
          <div className="card bg-primary text-white">
            <div className="card-body">
              <h3>Income</h3>
              <span>{convertToRupiah(income)}</span>
            </div>
          </div>
        </div>
        <div className="col-md-4 mt-4">
          <div className="card bg-danger text-white">
            <div className="card-body">
              <h3>Spending</h3>
              <span>{convertToRupiah(spending)}</span>
            </div>
          </div>
        </div>
        <div className="col-md-4 mt-4">
          <div className="card bg-warning text-white">
            <div className="card-body">
              <h3>Difference</h3>
              <span>{convertToRupiah(difference)}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Summary;
