import { HamburgerIcon, StarIcon } from "@chakra-ui/icons";
import { Flex, IconButton } from "@chakra-ui/react";
import { useAppDispatch, useAppSelector } from "@/store";
import { toggleSideBar } from "@/store/slices/sideBarSlice";
import NavItem from "./NavItem";

interface SideBarProps {
  bg: string;
  fg: string;
}

export default function SideBar({ bg, fg }: SideBarProps) {
  const dispatch = useAppDispatch();
  const { isCollapsed } = useAppSelector((state) => state.sideBar);

  return (
    <Flex
      position="relative"
      h="100vh"
      w={isCollapsed ? "65px" : "200px"}
      bg={bg}
      color={fg}
      flexShrink="0"
      boxShadow="md"
      transition="width 300ms ease"
    >
      <Flex
        p="3"
        flexDirection="column"
        alignItems="flex-start"
        gap="5"
        as="nav"
        w="full"
      >
        <IconButton
          aria-label="Toggle SideBar"
          background="none"
          icon={<HamburgerIcon />}
          onClick={() => dispatch(toggleSideBar())}
        />

        <NavItem questionNo={0} icon={StarIcon} title="Challenge 1" />
        <NavItem questionNo={1} icon={StarIcon} title="Challenge 2" />
        <NavItem questionNo={2} icon={StarIcon} title="Challenge 3" />
        <NavItem questionNo={3} icon={StarIcon} title="Challenge 4" />
        <NavItem questionNo={4} icon={StarIcon} title="Challenge 5" />
        <NavItem questionNo={5} icon={StarIcon} title="Challenge 6" />
      </Flex>
    </Flex>
  );
}
