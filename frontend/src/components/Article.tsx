import { type Article } from "@/types";
import defaultArticleImage from "@assets/default_article.jpg";
import { Link } from "react-router";

export const ArticleCard = ({ article }: Article) => {
  return (
    <Link
      to={"/article/" + article.id}
      className="card bg-gray-800 w-96 shadow-sm m-4 hover:bg-gray-600 transition-all ease-linear duration-75 hover:shadow-xl"
    >
      <figure>
        <img src={defaultArticleImage} alt="Article default image" />
      </figure>

      <div className="card-body rounded-b-xl">
        <h2 className="card-title">{article.title}</h2>
        <div className="flex">
          <p>{article.author}</p>
          <p className="-mx-0.5">@</p>
          <p className="text-gray-400">
            {new Date(article.created_at).toLocaleString()}
          </p>
        </div>
        <p>{article.description}</p>
      </div>
    </Link>
  );
};
