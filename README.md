# MVC

## Models
Es una entidad de tu BD -> [POSTS] -> POST

~~~javascript
class Post {
  constructor() {

  }

  getPosts() {
    // consulta de BD o una API
  }
}
~~~

## Controllers
Lógica del negocio
- Peticion del usuario -> Envia a otro lado

~~~php
class PostController {
  // /posts POST
  function createPost($request) {
    DD($request)
    Post::createPost($request);
  }
}

class PostRepository {
  function  createPost() {
    // 50 lineas
    return $post;
  }
}
~~~

## Views
Interfaz para el usuario, solamente se muestran los datos procesados.

~~~php
class PostController {
  // /posts GET
  function getPosts() {
    $posts = Post::all(); // select * from posts;
    return json_encode($posts);
  }

  // /posts POST
  function createPosts($request) {
    $post = Post::create($request); // select * from posts;


    // 201

    return json_encode($post);
  }
}

Route::get('/posts', PostCotnroller@index);
Route::post('/posts', PostCotnroller@create);
Route::put('/posts', PostCotnroller@create);
Route::patch('/posts', PostCotnroller@create);
~~~

fetch('lcoalhost/posts').then(darta => {})
fetch('lcoalhost/posts', {method: 'POST', body: JSON.strinfy(data)}).then(darta => {})
