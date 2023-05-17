const table = document.getElementById("users-table");
const form = document.querySelector("form");

const addRowsToTable = (data) => {
  for (let i = 0; i < data.length; i++) {
    tableRow(data[i]);
  }
};

const tableRow = (rowData) => {
  let row = table.insertRow(-1);
  row.id = rowData.Email;
  let cell1 = row.insertCell(0);
  let cell2 = row.insertCell(1);
  cell2.id = `${rowData.Email}-firstName`;
  let cell3 = row.insertCell(2);
  cell3.id = `${rowData.Email}-lastName`;
  let cell4 = row.insertCell(3);
  cell1.innerHTML = rowData.Email;
  cell2.innerHTML = rowData.FirstName;
  cell3.innerHTML = rowData.LastName;
  cell4.appendChild(
    addButtonToHtml(
      "btn btn-warning m-2",
      () => toggleEditUserModal({ ...rowData }),
      "Edit"
    )
  );
  cell4.appendChild(
    addButtonToHtml(
      "btn btn-danger",
      () => deleteRecord(rowData.Email),
      "Delete"
    )
  );
};

const updateUserNames = (data) => {
  let email = data.Email;
  let firstName = data.FirstName;
  let lastName = data.LastName;
  document.getElementById(`${email}-firstName`).innerHTML = firstName;
  document.getElementById(`${email}-lastName`).innerHTML = lastName;
};

const removeDataFromHtml = (email) => {
  let row = document.getElementById(email);
  row.remove();
};

const addButtonToHtml = (className, func, buttonText) => {
  let button = document.createElement("button");
  button.className = className;
  button.innerHTML = buttonText;
  button.onclick = func;
  return button;
};

const getDataAxios = () => {
  axios
    .get(API_GATEWAY_URL, {
      method: "GET",
      mode: "no-cors",
    })
    .then((response) => {
      addRowsToTable(response.data);
    })
    .catch((error) => {
      console.log(error);
    });
};

const deleteRecord = (email) => {
  axios
    .delete(`${API_GATEWAY_URL}?email=${email}`, {
      method: "DELETE",
      mode: "no-cors",
    })
    .then((response) => {
      removeDataFromHtml(email);
    })
    .catch((error) => {
      console.log(error);
    });
};

const createRecord = (event) => {
  event.preventDefault();

  let email = document.getElementById("user-email").value;
  let firstName = document.getElementById("user-first-name").value;
  let lastName = document.getElementById("user-last-name").value;
  let body = {
    email: email,
    firstName: firstName,
    lastName: lastName,
  };

  axios
    .post(API_GATEWAY_URL, body)
    .then((response) => {
      tableRow(response.data);
    })
    .catch((error) => {
      console.log(error);
    });

  toggleModal();
};

const toggleEditUserModal = (rowData) => {
  document.getElementById("user-email").value = rowData.Email;
  document.getElementById("user-email").disabled = true;

  document.getElementById("user-first-name").value = rowData.FirstName;
  document.getElementById("user-last-name").value = rowData.LastName;

  document.getElementById("modal-title").innerHTML = "Edit user";

  form.removeEventListener("submit", (e) => createRecord(e));
  form.addEventListener("submit", (e) => updateRecord(e));
  toggleModal();
};

const updateRecord = (event) => {
  event.preventDefault();

  let email = document.getElementById("user-email").value;
  let firstName = document.getElementById("user-first-name").value;
  let lastName = document.getElementById("user-last-name").value;
  let body = {
    email: email,
    firstName: firstName,
    lastName: lastName,
  };

  axios
    .put(API_GATEWAY_URL, body)
    .then((response) => {
      updateUserNames(response.data);
    })
    .catch((error) => {
      console.log(error);
    });
  toggleModal();
};

function toggleModal() {
  const userModal = document.getElementById("exampleModal");
  let modal = bootstrap.Modal.getInstance(userModal);
  if (modal) {
    modal.toggle();
  } else {
    let myModal = new bootstrap.Modal(userModal, {
      keyboard: false,
    });
    myModal.toggle();
  }
}

const toggleCreateUserModal = () => {
  document.getElementById("user-email").value = "";
  document.getElementById("user-email").disabled = false;

  document.getElementById("user-first-name").value = "";
  document.getElementById("user-last-name").value = "";

  form.removeEventListener("submit", (e) => updateRecord(e));
  form.addEventListener("submit", (e) => createRecord(e));

  document.getElementById("modal-title").innerHTML = "Create new user";

  toggleModal();
};

window.addEventListener("DOMContentLoaded", (event) => {
  getDataAxios();
});
