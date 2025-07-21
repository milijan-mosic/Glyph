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

  return (
    <div className="flex flex-row justify-center items-center">
      <button onClick={() => getArticle()}>Get</button>
    </div>
  );
}

export default Homepage;
