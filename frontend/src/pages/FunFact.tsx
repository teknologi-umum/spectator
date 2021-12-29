import Layout from "@/components/Layout";
import { withFinal } from "@/hoc";
import { Heading } from "@chakra-ui/react";

function FunFact() {
  return (
    <Layout>
      <Heading>
        Atoms are 99% nothing which means you&apos;re 99% nothing
      </Heading>
    </Layout>
  );
}

export default withFinal(FunFact);
