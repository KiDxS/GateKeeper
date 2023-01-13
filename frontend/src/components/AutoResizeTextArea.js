import React from "react";
import ResizeTextArea from "react-textarea-autosize";
import { Textarea } from "@chakra-ui/react";

const AutoResizeTextArea = React.forwardRef((props, ref) => {
    return (
        <Textarea
            size="lg"
            maxH="400px"
            ref={ref}
            minRows={1}
            as={ResizeTextArea}
            {...props}
            resize="none"
            isDisabled
        />
    );
});

export default AutoResizeTextArea;
