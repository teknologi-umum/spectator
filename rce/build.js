import esbuild from "esbuild";

esbuild
  .build({
    entryPoints: ["src/index.ts"],
    bundle: true,
    format: "esm",
    platform: "node",
    external: ["@grpc/grpc-js", "@protobuf-ts/runtime-rpc"],
    outdir: "./dist",
    target: ["es2021", "node16.14"],
    tsconfig: "tsconfig.json"
  })
  .catch((e) => {
    /* eslint-disable-next-line */
    console.error(e);
    process.exit(1);
  });
