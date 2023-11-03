1. Handler Layer
- Handle incoming HTTP requests;
- Validate input;
- Propagate calls to the service layer.

2. Service Layer
- Implement business logic;
- Return well-structured error types;
- Log errors.

3. Repository Layer
- Database interaction;
- Handle database-related errors;
- Log database-related errors.

### Routing Requessts
|**HTTP Method**|**URL Pattern**|**Handler**|**Action Description**|
|---|---|---|---|
|GET|/|getAllPosts|Show the list of all posts in JSON format|
|GET|/posts?post_id=1|getPostById|Show the post with post_id=1|
|POST|/posts/create|createPost|Create new post|
|POST|/posts/update?post_id=1|updatePost|Update the post with post_id=1|
|POST|/posts/delete?post_id=1|deletePost|Delete the post with post_id=1|
|#|#|#|#|
|GET|/categories/all|getAllCategories|Show the list of all categories|
|GET|/categories?category_id=1|getCategoryById|Show the category with category_id=1|
|POST|/categories/create|createCategory|Create new category|
|POST|/categories/update?category_id=1|updateCategory|Update the category with category_id=1|
|POST|/categories/delete?category_id=1|deleteCategory|Delete the category with category_id=1|
|#|#|#|#|
