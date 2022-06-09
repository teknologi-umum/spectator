import React from "react";
import {
  Modal,
  ModalBody,
  ModalContent,
  ModalHeader,
  ModalOverlay,
  Spinner
} from "@chakra-ui/react";

interface FinishedModalProps {
  isOpen: boolean;
}
export default function FinishedModal({ isOpen }: FinishedModalProps) {
  return (
    <>
      <Modal
        isOpen={isOpen}
        onClose={() => void 0}
        closeOnEsc={false}
        isCentered
      >
        <ModalOverlay />
        <ModalContent>
          <ModalHeader textAlign="center">Finishing Coding Test</ModalHeader>
          <ModalBody
            display="flex"
            alignItems="center"
            justifyContent="center"
            p={14}
          >
            <Spinner
              size="xl"
              thickness="4px"
              emptyColor="gray.200"
              color="blue.500"
            />
          </ModalBody>
        </ModalContent>
      </Modal>
    </>
  );
}
