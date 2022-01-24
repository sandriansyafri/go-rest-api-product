# BRIEF TEST_SANDRIAN SYAFRI_BACKEND DEVEPLOER

## Init
- first, you can import file go_api-project_product.sql to DB 
- next, set GOPATH=YourDirectory/go-rest-api-product


<hr/>

## Run App
- first,  go run main.go


<hr/>

## API

### Product 
<table>
  <thead>
    <tr>
      <th>#</th>
      <th>Method</th>
      <th>Description</th>
      <th>URL / Endpoint</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>1</td>
      <td>GET</td>
      <td>get all product</td>
      <td>localhost:8000/api/v1/products</td>
    </tr>
      <tr>
      <td>2</td>
      <td>GET</td>
      <td>get product by id</td>
      <td>localhost:8000/api/v1/products/:id</td>
    </tr>
     <tr>
      <td>3</td>
      <td>GET</td>
      <td>get product by name</td>
      <td>localhost:8000/api/v1/products?s=keywords</td>
    </tr>
    <tr>
      <td>3</td>
       <td>POST</td>
      <td>create product</td>
      <td>localhost:8000/api/v1/products</td>
    </tr>
    <tr>
      <td>4</td>
      <td>PUT</td>
      <td>update product</td>
       <td>localhost:8000/api/v1/products/:id</td>
    </tr>
     <tr>
      <td>5</td>
       <td>DELETE</td>
      <td>delete product</td>
       <td>localhost:8000/api/v1/products/:id</td>
    </tr>
  </tbody>
</table>

