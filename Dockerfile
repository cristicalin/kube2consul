FROM frolvlad/alpine-glibc:alpine-3.6

ENV K2C_CONSUL_API "127.0.0.1:8500"
ENV K2C_TOKEN_API ""
ENV K2C_KUBERNETES_API ""
ENV K2C_KUBECONFIG ""
ENV K2C_RESYNC_PERIOD "30"
ENV LOCK ""

WORKDIR /
COPY bin/kube2consul /kube2consul
EXPOSE 8080

CMD /kube2consul -logtostderr ${LOCK:+-lock}
