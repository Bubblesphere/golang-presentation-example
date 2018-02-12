const endpoints = {
  create: 'http://localhost:8000/task/create',
  list: 'http://localhost:8000/task/list',
  delete: 'http://localhost:8000/task/delete'
}

const createAPI = (content) => {  
  fetch(endpoints.create,  {
    method: 'post',
    headers: { 'Content-Type': 'text/plain' },
    body: content
  })
  .then(response => response.text())
  .then(json => addTodoToDom(json, content));
}

const listAPI = () => {  
  fetch(endpoints.list,  {
      method: 'post',
      headers: { 'Content-Type': 'text/plain' }
    })
    .then(response => response.text())
    .then(json => addListToDom(JSON.parse(json)));
}

const deleteAPI = (id) => {
  fetch(endpoints.delete,  {
    method: 'post',
    headers: { 'Content-Type': 'text/plain' },
    body: id
  });
}

const addListToDom = (json) => {  
  console.log("test")
  Object.keys(json).forEach(function(key) {
    addTodoToDom(key, json[key]);
  });
}

const addTodoToDom = (uuid, content) => {
  const li = createLi(uuid, content);
  const del = createDelete();
  del.addEventListener('click', (e) => {
    li.remove();
    deleteAPI(uuid);
  });
  li.appendChild(del);
  document.querySelectorAll('#todos')[0].appendChild(li);
}

const createLi = (uuid, content) => {
  let li = document.createElement("li");
  li.textContent = content;
  li.setAttribute("id", uuid);
  return li;
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
},false);

listAPI();