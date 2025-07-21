import axios from "axios";

function Homepage() {
  const url: string = "/api/1.0/article/";

  const getArticle = () => {
    axios
      .get(url + "get")
      .then(function (response) {
        console.log(response?.data);
      })
      .catch(function (error) {
        console.log(error);
      });
  };

  const createArticle = () => {
    axios
      .post(url + "create")
      .then(function (response) {
        console.log(response?.data);
      })
      .catch(function (error) {
        console.log(error);
      });
  };

  const updateArticle = () => {
    axios
      .put(url + "update")
      .then(function (response) {
        console.log(response?.data);
      })
      .catch(function (error) {
        console.log(error);
      });
  };

  const deleteArticle = () => {
    axios
      .delete(url + "delete")
      .then(function (response) {
        console.log(response?.data);
      })
      .catch(function (error) {
        console.log(error);
      });
  };

  return (
    <div className="flex flex-row justify-center items-center">
      <button onClick={() => getArticle()}>Get</button>
      <button onClick={() => createArticle()}>Create</button>
      <button onClick={() => updateArticle()}>Update</button>
      <button onClick={() => deleteArticle()}>Delete</button>
    </div>
  );
}

export default Homepage;
