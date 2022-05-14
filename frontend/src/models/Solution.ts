import { LRParser } from "@lezer/lr";
import { Language as LanguageEnum } from "@/stub/enums";
import { parser as javascriptParser } from "@lezer/javascript";
import { parser as phpParser } from "@lezer/php";
import { parser as javaParser } from "@lezer/java";
import { parser as cppParser } from "@lezer/cpp";
import { parser as pythonParser } from "@lezer/python";
import type { Language } from "@/models/Language";

export class Solution {
  private readonly LANGUAGE_TO_ENUM: Record<Language, LanguageEnum> = {
    c: LanguageEnum.C,
    cpp: LanguageEnum.CPP,
    java: LanguageEnum.JAVA,
    javascript: LanguageEnum.JAVASCRIPT,
    php: LanguageEnum.PHP,
    python: LanguageEnum.PYTHON
  };

  private readonly _languageParser = {
    [LanguageEnum.UNDEFINED]: undefined,
    [LanguageEnum.C]: cppParser,
    [LanguageEnum.CPP]: cppParser,
    [LanguageEnum.PHP]: phpParser,
    [LanguageEnum.JAVASCRIPT]: javascriptParser,
    [LanguageEnum.JAVA]: javaParser,
    [LanguageEnum.PYTHON]: pythonParser
  };

  private readonly _languageDirectiveType = {
    [LanguageEnum.UNDEFINED]: undefined,
    [LanguageEnum.C]: "PreprocDirective",
    [LanguageEnum.CPP]: "PreprocDirective",
    [LanguageEnum.PHP]: undefined,
    [LanguageEnum.JAVASCRIPT]: "ImportDeclaration",
    [LanguageEnum.JAVA]: "ImportDeclaration",
    [LanguageEnum.PYTHON]: "ImportStatement"
  };

  private readonly _content: string;
  private readonly _lang: LanguageEnum;
  private readonly _parser: LRParser;
  private readonly _directiveNodeType: string;

  constructor(lang: Language, content: string) {
    if (content === null || content === undefined) {
      throw new Error("Solution must be defined");
    }

    if (lang === null || lang === undefined) {
      throw new Error("Language must be defined");
    }

    const languageEnum = this.LANGUAGE_TO_ENUM[lang];
    const parser = this._languageParser[languageEnum];
    const directiveNodeType = this._languageDirectiveType[languageEnum];

    if (parser === undefined || directiveNodeType === undefined) {
      throw new Error(`Language ${lang} is not supported`);
    }

    this._content = content;
    this._lang = languageEnum;
    this._parser = parser;
    this._directiveNodeType = directiveNodeType;
  }

  public get language() {
    return this._lang;
  }

  public get content() {
    return this._content;
  }

  public getDirective() {
    const tree = this._parser.parse(this._content);
    return tree.topNode
      .getChildren(this._directiveNodeType)
      .map((b) => this._content.slice(b.from, b.to))
      .filter((directive) => {
        // C/C++ special case
        // filter out any preproc directive that isn't being used to include a header file
        if (
          this._directiveNodeType === "PreprocDirective" &&
          !directive.startsWith("#include")
        ) {
          return false;
        }

        return true;
      })
      .join("\n");
  }
}
