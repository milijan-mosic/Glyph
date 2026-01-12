import axios from "axios";
import { useEffect, useState } from "react";
import type { Comment } from "@/types";

export const AdminCommentsPage = () => {
  const [comments, setComments] = useState<Comment[]>([]);
  const [loading, setLoading] = useState<boolean>(true);

  const fetchPendingComments = async () => {
    try {
      const res = await axios.get("/api/1.0/comment/pending");
      setComments(res.data);
    } catch (err) {
      console.error("Failed to fetch pending comments", err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchPendingComments();
  }, []);

  const approveComment = async (id: number) => {
    try {
      await axios.put(`/api/1.0/comment/approve/${id}`);
      setComments((prev) => prev.filter((c) => c.id !== id));
    } catch (err) {
      console.error("Failed to approve comment", err);
    }
  };

  const deleteComment = async (id: number) => {
    try {
      await axios.delete(`/api/1.0/comment/delete/${id}`);
      setComments((prev) => prev.filter((c) => c.id !== id));
    } catch (err) {
      console.error("Failed to delete comment", err);
    }
  };

  if (loading) {
    return <p className="text-gray-400">Loading pending comments…</p>;
  }

  if (comments.length === 0) {
    return <p className="text-gray-400">No pending comments 🎉</p>;
  }

  return (
    <div className="flex justify-center" data-testid="admin-comments-page">
      <div className="w-[800px] p-8 bg-gray-800 rounded-xl">
        <h1 className="text-2xl font-bold mb-6">Pending Comments</h1>

        <div className="space-y-4" data-testid="comments-section">
          {comments.map((comment) => (
            <div
              key={comment.id}
              className="p-4 rounded-lg bg-gray-700"
              data-testid="admin-comment" // generalized for all comments
            >
              <p
                className="text-sm text-gray-300"
                data-testid="comment-article-id"
              >
                Article ID: {comment.post_id} ·{" "}
                {new Date(comment.created_at).toLocaleString()}
              </p>

              <p className="font-semibold mt-1" data-testid="comment-author">
                {comment.author_name}
              </p>

              <p
                className="mt-2 whitespace-pre-wrap"
                data-testid="comment-content"
              >
                {comment.content}
              </p>

              <div className="mt-4 flex gap-2">
                <button
                  className="btn btn-success btn-sm approve-btn"
                  data-testid="approve-btn"
                  onClick={() => approveComment(comment.id)}
                >
                  Approve
                </button>

                <button
                  className="btn btn-error btn-sm"
                  data-testid="delete-btn"
                  onClick={() => deleteComment(comment.id)}
                >
                  Delete
                </button>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};
