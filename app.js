const endpoints = {
  create: 'http://localhost:8000/task/create',
  list: 'http://localhost:8000/task/list',
  delete: 'http://localhost:8000/task/delete'
}

const createAPI = (content) => {
  fetch(endpoints.create, {
    method: 'post',
    headers: { 'Content-Type': 'text/plain' },
    body: content
  })
    .then(response => response.text())
    .then(json => addTodoToDom(json, content));
}

const listAPI = () => {
  fetch(endpoints.list, {
    method: 'post',
    headers: { 'Content-Type': 'text/plain' }
  })
    .then(response => response.text())
    .then(json => addListToDom(JSON.parse(json)));
}

const deleteAPI = (id) => {
  fetch(endpoints.delete, {
    method: 'post',
    headers: { 'Content-Type': 'text/plain' },
    body: id
  });
}

const addListToDom = (json) => {
  Object.keys(json).forEach(function (key) {
    addTodoToDom(key, json[key]);
  });
}

const addTodoToDom = (uuid, content) => {
  const li = createRow(uuid, content);
  document.querySelectorAll('#todos')[0].appendChild(li);
}

const createRow = (uuid, content) => {
  let row = document.createElement("div");
  let txt = document.createElement("div");
  let btn = document.createElement("div");

  const del = createDelete();
  del.className = "waves-effect waves-light btn red";
  del.addEventListener('click', (e) => {
    row.remove();
    deleteAPI(uuid);
  });

  row.className = "row";
  txt.className = "col s11";
  txt.textContent = content;

  btn.className = "col s1"; 
  btn.appendChild(del);
  btn.setAttribute("id", uuid);
  
  row.appendChild(txt);
  row.appendChild(btn);

  return row;
}

const createDelete = () => {
  let del = document.createElement("input");
  del.value = 'x';
  del.type = 'button';
  return del;
}

const add = document.querySelectorAll('#add')[0];
add.addEventListener("click", (e) => {
  const content = document.querySelectorAll('#content')[0];
  if (content.value !== '') {
    createAPI(content.value);
    content.value = '';
  }
}, false);

listAPI();