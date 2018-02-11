const endpoints = {
  create: 'http://localhost:8000/task/create',
  list: 'http://localhost:8000/task/list',
  delete: 'http://localhost:8000/task/delete'
}

const createAPI = (content) => {
  addTodoToDom('0e879e24-9894-4dfa-73a7-031cd1dc2a21', content);
  /*fetch(endpoints.create,  {
    method: 'post',
    headers: { 'Content-Type': 'text/plain' },
    body: content
  })
  .then(response => response.json())
  .then(json => addTodoToDom({
    uuid: json,
    content: content
  }));*/
}

const listAPI = () => {
  addListToDom(JSON.parse('{"0e879e24-9894-4dfa-73a7-031cd1dc2a21":"1","32c98be7-25fe-42d7-7c93-f97e09db7a1d":"2","4fd7b310-4573-47d0-6b21-49ef614d1970":"3"}'));
  /*fetch(endpoints.list,  {
      method: 'post',
      headers: { 'Content-Type': 'text/plain' }
    })
    .then(response => response.json())
    .then(json => addListToDom(JSON.parse(json)));*/
}

const deleteAPI = (id) => {
  fetch(endpoints.delete,  {
    method: 'delete'
  });
}

const addListToDom = (json) => {
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