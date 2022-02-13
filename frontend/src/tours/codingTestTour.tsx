import React from "react";
import type { StepType } from "@reactour/tour";

export const codingTestTour: StepType[] = [
  {
    selector: "[data-tour=\"sidebar-step-1\"]",
    content: <p>You can press this button to toggle the sidebar.</p>,
    disableActions: true
  },
  {
    selector: "[data-tour=\"sidebar-step-2\"]",
    content: (
      <p>You can switch between questions by pressing one of these buttons.</p>
    ),
    disableActions: false
  },
  {
    selector: "[data-tour=\"topbar-step-1\"]",
    content: (
      <p>
        This is how much time you have left for the Coding Test. When it goes to
        zero, you will automatically get redirected to the next page.
      </p>
    )
  },
  {
    selector: "[data-tour=\"topbar-step-2\"]",
    content: (
      <p>
        You can change the theme by selecting one of the options in this
        dropdown.
      </p>
    )
  },
  {
    selector: "[data-tour=\"topbar-step-3\"]",
    content: (
      <p>
        You can select which programming language you want to use to solve the
        question by selecting one of the options in this dropdown.
      </p>
    )
  },
  {
    selector: "[data-tour=\"topbar-step-4\"]",
    content: (
      <p>
        If the font size is too small or too big, you can change it by selecting
        one of the options in this dropdown.
      </p>
    )
  },
  {
    selector: "[data-tour=\"topbar-step-5\"]",
    content: (
      <p>
        You can change the language of the User Interface by selecting one of
        the options in this dropdown.
      </p>
    )
  },
  {
    selector: "[data-tour=\"topbar-step-6\"]",
    content: (
      <p>
        If you don&apos;t think you can solve the question anymore, you can
        press this button. Please don&apos;t use it unless you&apos;re really
        sure.
      </p>
    )
  },
  {
    selector: "[data-tour=\"topbar-step-7\"]",
    content: (
      <p>
        You can press this button to test your answer. This won&apos;t submit
        your answer so you can press it multiple times to check the correctness
        of your answer.
      </p>
    )
  },
  {
    selector: "[data-tour=\"topbar-step-8\"]",
    content: (
      <p>
        You can press this button to submit your answer. I&apos;d recommend
        testing your answer multiple times using the Test button previously
        before trying to submit it.
      </p>
    )
  },
  {
    selector: "[data-tour=\"question-step-1\"]",
    content: (
      <p>
        This is the Prompt tab. Every detail you need to solve the question is
        laid out in this tab.
      </p>
    )
  },
  {
    selector: "[data-tour=\"question-step-2\"]",
    content: (
      <p>
        This is the Result tab. Grayed out means that you haven&apos;t run your
        code. The result of your submission can be seen here.
      </p>
    )
  },
  {
    selector: "[data-tour=\"editor-step-1\"]",
    content: (
      <>
        <p>
          This is your code editor. You can write your code here. Please
          don&apos;t change the boilerplate function, otherwise your submission
          won&apos;t be accepted.
        </p>
        <br />
        <p>
          Your code will be persisted in the browser&apos;s local storage. You
          can switch between questions, refresh the page, restart the browser,
          and your code will still be here.
        </p>
      </>
    )
  },
  {
    selector: "[data-tour=\"scratchpad-step-1\"]",
    content: (
      <>
        <p>
          This is your scratchpad. You can write anything on it. It&apos;s just
          there to help you, it won&apos;t be submitted.
        </p>
        <br />
        <p>
          The content of the scratchpad will be persisted in the browser&apos;s
          local storage as well so don&apos;t worry about losing it.
        </p>
      </>
    )
  }
];
