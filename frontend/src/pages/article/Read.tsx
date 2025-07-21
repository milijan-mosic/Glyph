import axios from "axios";
import { useEffect, useState } from "react";
import { type Article } from "@/types";
import { useParams } from "react-router";
import defaultArticleImage from "@assets/default_article.jpg";

export const ReadArticle = () => {
  const { articleId } = useParams<string>();
  const [article, setArticle] = useState<Article>({});

  const fetchArticle = async () => {
    const url: string = "/api/1.0/article/";

    const article: Article = await axios
      .get(url + "get/" + articleId)
      .then((response) => {
        return response?.data?.article;
      })
      .catch((error) => {
        console.log(error);
        return {};
      });

    setArticle(article);
  };

  useEffect(() => {
    fetchArticle();
  }, [articleId]);

  return (
    <div className="flex flex-col justify-center items-center">
      <div className="w-[700px]">
        <img
          src={defaultArticleImage}
          alt="Article default image"
          className="rounded-b-xl"
        />

        <div className="flex flex-row my-4">
          <p>{article.author}</p>
          <p className="mx-2">@</p>
          <p className="text-gray-400">{article.created_at}</p>
        </div>

        <h2 className="text-3xl">{article.title}</h2>

        <div className="mt-8">{article.content}</div>
      </div>
    </div>
  );
};

//  <button className="btn btn-warning">Edit</button>
