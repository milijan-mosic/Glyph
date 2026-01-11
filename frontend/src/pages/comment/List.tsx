import axios from "axios";
import { useEffect, useState } from "react";
import type { Comment } from "@/types";

type Props = {
  articleId: string;
};

export const CommentList = ({ articleId }: Props) => {
  const [comments, setComments] = useState<Comment[]>([]);
  const [loading, setLoading] = useState<boolean>(true);

  const fetchComments = async () => {
    try {
      const res = await axios.get(
        `/api/1.0/comment/list/${articleId}`
      );
      setComments(res.data);
    } catch (err) {
      console.error("Failed to fetch comments", err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchComments();
  }, [articleId]);

  if (loading) {
    return <p className="text-gray-400 mt-6">Loading comments…</p>;
  }

  if (comments.length === 0) {
    return <p className="text-gray-400 mt-6">No comments yet.</p>;
  }

  return (
    <div className="mt-8 space-y-4">
      <h3 className="text-xl font-semibold">Comments</h3>

      {comments.map((comment) => (
        <div
          key={comment.id}
          className="p-4 rounded-lg bg-gray-700"
        >
          <p className="text-sm text-gray-300">
            {comment.author_name} ·{" "}
            {new Date(comment.created_at).toLocaleString()}
          </p>

          <p className="mt-2 whitespace-pre-wrap">
            {comment.content}
          </p>
        </div>
      ))}
    </div>
  );
};
