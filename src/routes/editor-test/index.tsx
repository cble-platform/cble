import Editor from "@monaco-editor/react";
import { useEffect, useState } from "react";

export default function EditorTest() {
  const [editorValue, setEditorValue] = useState<string>("");

  useEffect(() => {
    console.log(editorValue);
  }, [editorValue]);

  return (
    <>
      <Editor
        height="90vh"
        defaultLanguage="javascript"
        defaultValue="// some comment"
        onChange={(value, _) => setEditorValue(value ?? "")}
      />
    </>
  );
}
