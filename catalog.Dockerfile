FROM scratch

LABEL operators.operatorframework.io.index.configs.v1=/configs
LABEL operators.operatorframework.io.index.mediatype.v1=registry+v1

COPY catalog /configs
