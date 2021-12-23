import { Box, Button, useColorMode } from "@chakra-ui/react";


export default function ThemeButton() {
  const { toggleColorMode } = useColorMode()
  return (
    <Button position="fixed" onClick={toggleColorMode}>
      Toggle mode
    </Button>
  )
}