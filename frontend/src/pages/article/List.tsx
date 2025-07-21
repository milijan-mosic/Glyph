import axios from "axios";
import { useEffect, useState } from "react";
import { type Article } from "@/types";
import { ArticleCard } from "@components/Article";

function Homepage() {
  const fetchArticles = async () => {
    const url: string = "/api/1.0/article/";

    const articles: Array<Article> = await axios
      .get(url + "list")
      .then((response) => {
        return response?.data?.articles;
      })
      .catch((error) => {
        console.log(error);
        return [];
      });

    setArticles(articles);
  };

  const [articles, setArticles] = useState<Array<Article>>([]);

  useEffect(() => {
    fetchArticles();
  }, []);

  return (
    <div className="flex flex-row justify-center">
      <div>
        {articles.map((article: Article) => (
          <ArticleCard key={article.id} article={article} />
        ))}
      </div>
    </div>
  );
}

export default Homepage;
