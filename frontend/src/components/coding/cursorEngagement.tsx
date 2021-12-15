export const mouseOverFunc = (str, isCode, cursorOverSections, highEngagementSections, middleEngagementSections, lowEngagementSections, enterTime) => {
    const duration = (new Date() - enterTime)/1000;

    console.log("duration ----------> ", duration);

    // question (Average reading speed 300 wpm)
    // if length >= 3, then > 0.6 sec
    // if length < 3, then > 0.4 sec

    // code (reading speed slightly slower than question reading 240 wpm)
    // if length >= 3, then > 0.75 sec
    // if length < 3, then > 0.5 sec

    const quesTimeEqual3High = 0.6 * (4/3);
    const quesTimeEqual3Low = 0.6 * (2/3);
    const quesTimeEqual2High = 0.4 * (4/3);
    const quesTimeEqual2Low = 0.4 * (2/3);

    const codeTimeEqual3High = 0.75 * (4/3);
    const codeTimeEqual3Low = 0.75 * (2/3);
    const codeTimeEqual2High = 0.5 * (4/3);
    const codeTimeEqual2Low = 0.5 * (2/3);

    if (!cursorOverSections.includes(str)) {
        cursorOverSections.push(str);
    }
    if (str.split(/[ ;"()=?]/).filter(l => l !== "").length < 3) {
        if (isCode === true) {
            if (duration >= codeTimeEqual2High) {
                if (!highEngagementSections.includes(str)) {
                    highEngagementSections.push(str);
                    if (lowEngagementSections.includes(str)) {
                        const index = lowEngagementSections.indexOf(str);
                        lowEngagementSections.splice(index, 1);
                    }
                    if (middleEngagementSections.includes(str)) {
                        const index = middleEngagementSections.indexOf(str);
                        middleEngagementSections.splice(index, 1);
                    }
                }
            } else if (duration <= codeTimeEqual2Low) {
                if (!highEngagementSections.includes(str) && !middleEngagementSections.includes(str)
                && !lowEngagementSections.includes(str)) {
                    lowEngagementSections.push(str);
                }
            } else {
                if (!highEngagementSections.includes(str) && !middleEngagementSections.includes(str)) {
                    middleEngagementSections.push(str);
                    console.log("middle entry");
                    if (lowEngagementSections.includes(str)) {
                        const index = lowEngagementSections.indexOf(str);
                        lowEngagementSections.splice(index, 1);
                    }
                }
            }
        } else {
            if (duration >= quesTimeEqual2High) {
                if (!highEngagementSections.includes(str)) {
                    highEngagementSections.push(str);
                    if (lowEngagementSections.includes(str)) {
                        const index = lowEngagementSections.indexOf(str);
                        lowEngagementSections.splice(index, 1);
                    }
                    if (middleEngagementSections.includes(str)) {
                        const index = middleEngagementSections.indexOf(str);
                        middleEngagementSections.splice(index, 1);
                    }
                }
            } else if (duration <= quesTimeEqual2Low) {
                if (!highEngagementSections.includes(str) && !middleEngagementSections.includes(str)
                && !lowEngagementSections.includes(str)) {
                    lowEngagementSections.push(str);
                }
            } else {
                if (!highEngagementSections.includes(str) && !middleEngagementSections.includes(str)) {
                    middleEngagementSections.push(str);
                    if (lowEngagementSections.includes(str)) {
                        const index = lowEngagementSections.indexOf(str);
                        lowEngagementSections.splice(index, 1);
                    }
                }
            }
        }
    } else {
        if (isCode === true) {
            if (duration >= codeTimeEqual3High) {
                if (!highEngagementSections.includes(str)) {
                    highEngagementSections.push(str);
                    if (lowEngagementSections.includes(str)) {
                        const index = lowEngagementSections.indexOf(str);
                        lowEngagementSections.splice(index, 1);
                    }
                    if (middleEngagementSections.includes(str)) {
                        const index = middleEngagementSections.indexOf(str);
                        middleEngagementSections.splice(index, 1);
                    }
                }
            } else if (duration <= codeTimeEqual3Low) {
                if (!highEngagementSections.includes(str) && !middleEngagementSections.includes(str)
                && !lowEngagementSections.includes(str)) {
                    lowEngagementSections.push(str);
                }
            } else {
                if (!highEngagementSections.includes(str) && !middleEngagementSections.includes(str)) {
                    middleEngagementSections.push(str);
                    if (lowEngagementSections.includes(str)) {
                        const index = lowEngagementSections.indexOf(str);
                        lowEngagementSections.splice(index, 1);
                    }
                }
            }
        } else {
            if (duration >= quesTimeEqual3High) {
                if (!highEngagementSections.includes(str)) {
                    highEngagementSections.push(str);
                    if (lowEngagementSections.includes(str)) {
                        const index = lowEngagementSections.indexOf(str);
                        lowEngagementSections.splice(index, 1);
                    }
                    if (middleEngagementSections.includes(str)) {
                        const index = middleEngagementSections.indexOf(str);
                        middleEngagementSections.splice(index, 1);
                    }
                }
            } else if (duration <= quesTimeEqual3Low) {
                if (!highEngagementSections.includes(str) && !middleEngagementSections.includes(str)
                && !lowEngagementSections.includes(str)) {
                    lowEngagementSections.push(str);
                }
            } else {
                if (!highEngagementSections.includes(str) && !middleEngagementSections.includes(str)) {
                    middleEngagementSections.push(str);
                    if (lowEngagementSections.includes(str)) {
                        const index = lowEngagementSections.indexOf(str);
                        lowEngagementSections.splice(index, 1);
                    }
                }
            }
        }
    }
};

export const highlightedFunc = (str, highEngagementSections, middleEngagementSections, lowEngagementSections) => {
    if (!highEngagementSections.includes(str)) {
        highEngagementSections.push(str);
        if (lowEngagementSections.includes(str)) {
            const index = lowEngagementSections.indexOf(str);
            lowEngagementSections.splice(index, 1);
        }
        if (middleEngagementSections.includes(str)) {
            const index = middleEngagementSections.indexOf(str);
            middleEngagementSections.splice(index, 1);
        }
    }
};