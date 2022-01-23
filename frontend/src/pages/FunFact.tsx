import React from "react";
import Layout from "@/components/Layout";
import { Heading } from "@chakra-ui/react";
import { useEffect } from "react";

export default function FunFact() {
  useEffect(() => {
    document.title = "Fun Fact | Spectator";
  }, []);

  return (
    <Layout>
      <Heading>
        Atoms are 99% nothing which means you&apos;re 99% nothing
      </Heading>
    </Layout>
  );
}

