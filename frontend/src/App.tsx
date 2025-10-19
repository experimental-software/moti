import "./App.css";
import { Container } from "react-bootstrap";
import { useEffect, useState } from "react";
import { GetExampleParagraphs } from "../wailsjs/go/main/App";
import { splitParagraph } from "./model/text";

function App() {
  const [searchTerm, setSearchTerm] = useState<string>("");
  const [paragraphs, setParagraphs] = useState<string[]>([]);

  useEffect(() => {
    GetExampleParagraphs().then((p) => setParagraphs(p));
  }, []);

  return (
    <Container>
      <div className={"header"}>
        <h1>Document search</h1>
        <input
          className="form-control"
          type="text"
          placeholder="Search term"
          onChange={async (event) => {
            const searchTerm = event.target.value;
            setSearchTerm(searchTerm);
          }}
        />
      </div>
      <div className={"content"}>
        <div className={"preview"}>
          {paragraphs.map((value, index) => (
            <Paragraph
              searchTerm={searchTerm}
              text={value}
              key={`Paragraph${index}`}
            />
          ))}
        </div>
      </div>
    </Container>
  );
}

interface ParagraphProps {
  text: string;
  searchTerm: string;
}

function Paragraph(props: ParagraphProps) {
  const parts = splitParagraph(props.text, props.searchTerm);
  return (
    <div className={"paragraph"}>
      {parts.map((p, index) =>
        p.toLowerCase() === props.searchTerm.toLowerCase() ? (
          <mark key={`mark-${index}`}>{p}</mark>
        ) : (
          <span key={`span-${index}`}>{p}</span>
        ),
      )}
    </div>
  );
}

export default App;
