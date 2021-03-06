package mocks

//go:generate sh -c "mockgen -package mockquic -destination quic/stream.go github.com/yuyyi51/quic-go Stream && goimports -w quic/stream.go"
//go:generate sh -c "mockgen -package mockquic -destination quic/early_session_tmp.go github.com/yuyyi51/quic-go EarlySession && sed 's/qtls.ConnectionState/quic.ConnectionState/g' quic/early_session_tmp.go > quic/early_session.go && rm quic/early_session_tmp.go && goimports -w quic/early_session.go"
//go:generate sh -c "mockgen -package mockquic -destination quic/early_listener.go github.com/yuyyi51/quic-go EarlyListener && goimports -w quic/early_listener.go"
//go:generate sh -c "mockgen -package mocks -destination tracer.go github.com/yuyyi51/quic-go/logging Tracer && goimports -w tracer.go"
//go:generate sh -c "mockgen -package mocks -destination connection_tracer.go github.com/yuyyi51/quic-go/logging ConnectionTracer && goimports -w connection_tracer.go"
//go:generate sh -c "mockgen -package mocks -destination short_header_sealer.go github.com/yuyyi51/quic-go/internal/handshake ShortHeaderSealer && goimports -w short_header_sealer.go"
//go:generate sh -c "mockgen -package mocks -destination short_header_opener.go github.com/yuyyi51/quic-go/internal/handshake ShortHeaderOpener && goimports -w short_header_opener.go"
//go:generate sh -c "mockgen -package mocks -destination long_header_opener.go github.com/yuyyi51/quic-go/internal/handshake LongHeaderOpener && goimports -w long_header_opener.go"
//go:generate sh -c "mockgen -package mocks -destination crypto_setup_tmp.go github.com/yuyyi51/quic-go/internal/handshake CryptoSetup && sed 's/github.com\\/marten-seemann\\/qtls/github.com\\/lucas-clemente\\/quic-go\\/internal\\/qtls/g' crypto_setup_tmp.go > crypto_setup.go && rm crypto_setup_tmp.go && goimports -w crypto_setup.go"
//go:generate sh -c "mockgen -package mocks -destination stream_flow_controller.go github.com/yuyyi51/quic-go/internal/flowcontrol StreamFlowController && goimports -w stream_flow_controller.go"
//go:generate sh -c "mockgen -package mocks -destination congestion.go github.com/yuyyi51/quic-go/internal/congestion SendAlgorithmWithDebugInfos && goimports -w congestion.go"
//go:generate sh -c "mockgen -package mocks -destination connection_flow_controller.go github.com/yuyyi51/quic-go/internal/flowcontrol ConnectionFlowController && goimports -w connection_flow_controller.go"
//go:generate sh -c "mockgen -package mockackhandler -destination ackhandler/sent_packet_handler.go github.com/yuyyi51/quic-go/internal/ackhandler SentPacketHandler && goimports -w ackhandler/sent_packet_handler.go"
//go:generate sh -c "mockgen -package mockackhandler -destination ackhandler/received_packet_handler.go github.com/yuyyi51/quic-go/internal/ackhandler ReceivedPacketHandler && goimports -w ackhandler/received_packet_handler.go"

// The following command produces a warning message on OSX, however, it still generates the correct mock file.
// See https://github.com/golang/mock/issues/339 for details.
//go:generate sh -c "mockgen -package mocktls -destination tls/client_session_cache.go crypto/tls ClientSessionCache && goimports -w tls/client_session_cache.go"
