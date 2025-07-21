import axios from "axios";
import { useEffect, useState } from "react";
import { type Article } from "@/types";
import { ArticleCard } from "@components/Article";

export const ListArticles = () => {
  const url: string = "/api/1.0/article/";
  const [articles, setArticles] = useState<Array<Article>>([]);

  const fetchArticles = async () => {
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
};
