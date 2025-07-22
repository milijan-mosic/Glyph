import axios from "axios";
import {
  MDXEditor,
  headingsPlugin,
  listsPlugin,
  quotePlugin,
  thematicBreakPlugin,
  UndoRedo,
  BoldItalicUnderlineToggles,
  toolbarPlugin,
  linkDialogPlugin,
  imagePlugin,
  tablePlugin,
  BlockTypeSelect,
  CreateLink,
  InsertImage,
  ListsToggle,
  type MDXEditorMethods,
} from "@mdxeditor/editor";
import { type Article } from "@/types";

import "@mdxeditor/editor/style.css";
import { useEffect, useRef, useState } from "react";
import { useParams, useNavigate } from "react-router";

export const WriteOrEditArticle = () => {
  const url: string = "/api/1.0/article/";
  const editorRef = useRef<MDXEditorMethods>(null);

  const [title, setTitle] = useState<string>("");
  const [description, setDescription] = useState<string>("");
  const [published, setPublished] = useState<boolean>(false);
  const [content, setContent] = useState<string>("");

  const { articleId } = useParams<string>();
  const navigate = useNavigate();

  const fetchArticle = async () => {
    const article: Article = await axios
      .get(url + "get/" + articleId)
      .then((response) => {
        return response?.data?.article;
      })
      .catch((error) => {
        console.log(error);
        return {};
      });

    if (article) {
      setTitle(article?.title);
      setDescription(article?.description);
      setPublished(article?.published);
      setContent(article?.content);
    }
  };

  useEffect(() => {
    if (articleId !== "") {
      fetchArticle();
    }
  }, [articleId]);

  const createArticle = async () => {
    await axios
      .post(url + "create", {
        title: title,
        description: description,
        published: published,
        content: editorRef.current?.getMarkdown(),
        author: "MM",
      })
      .then((response) => {
        navigate("/article/" + response?.data?.ArticleId);
      })
      .catch((error) => {
        console.error("Axios error:", error.message);
      });
  };

  const updateArticle = async () => {
    await axios
      .put(url + "update", {
        article_id: articleId,
        title: title,
        description: description,
        published: published,
        content: editorRef.current?.getMarkdown(),
      })
      .then(() => {
        navigate("/article/" + articleId);
      })
      .catch((error) => {
        console.error("Axios error:", error.message);
      });
  };

  const deleteArticle = async () => {
    await axios
      .delete(url + "delete/" + articleId)
      .then(() => {
        navigate("/");
        document.getElementById("delete_article_modal").showModal();
      })
      .catch((error) => {
        console.error("Axios error:", error.message);
      });
  };

  const handleArticle = async () => {
    if (articleId) {
      updateArticle();
    } else {
      createArticle();
    }
  };

  return (
    <div className="flex flex-row justify-center items-center">
      <div className="flex flex-col w-[800px] p-8 rounded-b-xl bg-gray-800">
        <div className="flex justify-between">
          <fieldset className="fieldset">
            <legend className="fieldset-legend">Title</legend>
            <input
              type="text"
              className="input"
              placeholder="Type here"
              value={title}
              onChange={(e) => {
                setTitle(e.target.value);
              }}
            />
          </fieldset>

          <fieldset className="fieldset flex justify-end">
            <legend className="fieldset-legend">Published</legend>
            <input
              type="checkbox"
              checked={published}
              onChange={(e) => {
                setPublished(e.target.checked ? true : false);
              }}
              className="toggle border-red-600 bg-red-500 checked:border-green-500 checked:bg-green-400 checked:text-green-800"
            />
          </fieldset>
        </div>

        <fieldset className="fieldset mt-4">
          <legend className="fieldset-legend">Description</legend>
          <input
            type="text"
            className="input"
            placeholder="Type here"
            value={description}
            onChange={(e) => {
              setDescription(e.target.value);
            }}
          />
        </fieldset>

        <fieldset className="fieldset mt-4">
          <legend className="fieldset-legend">Content</legend>
          <MDXEditor
            ref={editorRef}
            className="bg-white rounded-xl w-[700px]"
            markdown={content}
            onChange={(e) => {
              setContent(e);
            }}
            plugins={[
              headingsPlugin(),
              listsPlugin(),
              quotePlugin(),
              thematicBreakPlugin(),
              linkDialogPlugin(),
              imagePlugin(),
              tablePlugin(),
              toolbarPlugin({
                toolbarContents: () => (
                  <>
                    <UndoRedo />
                    <BoldItalicUnderlineToggles />
                    <BlockTypeSelect />
                    <CreateLink />
                    <InsertImage />
                    <ListsToggle />
                  </>
                ),
              }),
            ]}
          />
        </fieldset>

        <div className="mt-8 flex justify-between">
          <button
            className="btn btn-error"
            onClick={() => {
              document.getElementById("delete_article_modal").showModal();
            }}
            disabled={articleId ? false : true}
          >
            Delete
          </button>

          <button className="btn btn-primary" onClick={() => handleArticle()}>
            {articleId ? "Edit" : "Create"}
          </button>
        </div>
      </div>

      <dialog id="delete_article_modal" className="modal">
        <div className="modal-box">
          <h3 className="font-bold text-lg">Delete article</h3>
          <p className="py-4">Are you sure?</p>

          <div className="modal-action">
            <form method="dialog">
              <button className="btn btn-neutral">Close</button>
            </form>
            <button className="btn btn-error" onClick={() => deleteArticle()}>
              Confirm
            </button>
          </div>
        </div>
      </dialog>
    </div>
  );
};
