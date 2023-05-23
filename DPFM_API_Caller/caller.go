package dpfmapicaller

import (
	"bufio"
	"context"
	dpfm_api_input_reader "data-platform-api-orders-edi-for-voluntary-chain-smes-csv-converter-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-orders-edi-for-voluntary-chain-smes-csv-converter-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-orders-edi-for-voluntary-chain-smes-csv-converter-rmq-kube/config"
	"io/ioutil"
	"os"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type DPFMAPICaller struct {
	ctx  context.Context
	conf *config.Conf
	log  *logger.Logger
}

func NewDPFMAPICaller(
	conf *config.Conf, l *logger.Logger,
) *DPFMAPICaller {
	return &DPFMAPICaller{
		ctx:  context.Background(),
		conf: conf,
		log:  l,
	}
}

func (c *DPFMAPICaller) OrdersEdiForVoluntaryChainSmes(rmq *rabbitmq.RabbitmqClient, reader *dpfm_api_input_reader.FileReader, f *dpfm_api_input_reader.Request) error {
	err := deleteSecondRow(f.Orders.FilePath)
	if err != nil {
		err = xerrors.Errorf("delete second row error: %w", err)
		data := c.createHeaderData(f, err)
		err = rmq.Send("data-platform-api-data-concatenation-queue", data)
		if err != nil {
			c.log.Error("data send error: %w", err)
		} else {
			c.log.JsonParseOut(data)
		}
		return nil
	}

	ordersMessage, err := dpfm_api_output_formatter.ConvertToConcatMessage(f.Orders.FilePath)
	if err != nil {
		return xerrors.Errorf("file convert error: %w", err)

	}

	for _, v := range *ordersMessage {
		data := c.setHeaderData(&v, f, err)
		err = rmq.Send("data-platform-api-data-concatenation-queue", data)
		if err != nil {
			c.log.Error("data send error: %w", err)
		} else {
			c.log.JsonParseOut(data)
		}
	}
	return nil
}

func deleteSecondRow(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return xerrors.Errorf("file open error: %w", err)
	}

	scanner := bufio.NewScanner(f)

	// 1行目コピー
	if !scanner.Scan() {
		f.Close()
		return xerrors.Errorf("ファイルの1行目の取得に失敗しました。")
	}
	data := []byte(scanner.Text())

	// 2行目は飛ばす
	if !scanner.Scan() {
		f.Close()
		return xerrors.Errorf("ファイルの2行目の取得に失敗しました。")
	}

	// 3行目以降
	for scanner.Scan() {
		data = append(data, '\n')
		data = append(data, []byte(scanner.Text())...)
	}

	f.Close()
	err = os.Remove(filePath)
	if err != nil {
		return xerrors.Errorf("file close error: %w", err)
	}

	err = ioutil.WriteFile(filePath, []byte(data), 0664)
	if err != nil {
		return xerrors.Errorf("file write error: %w", err)
	}
	return nil
}

func (c *DPFMAPICaller) setHeaderData(ordersMessage *dpfm_api_output_formatter.DataConcatenation, f *dpfm_api_input_reader.Request, apiErr error) *dpfm_api_output_formatter.OrdersSDC {
	var apiProcessingError *string = nil
	if apiErr != nil {
		e := apiErr.Error()
		apiProcessingError = &e
	}

	return &dpfm_api_output_formatter.OrdersSDC{
		ConnectionKey:       f.ConnectionKey,
		Result:              apiErr == nil,
		RedisKey:            f.RedisKey,
		Filepath:            f.Filepath,
		APIStatusCode:       f.APIStatusCode,
		RuntimeSessionID:    f.RuntimeSessionID,
		BusinessPartnerID:   f.BusinessPartnerID,
		ServiceLabel:        "FUNCTION_ORDERS_EDI_FOR_VOLUNTARY_CHAIN_SMES_DATA_CONCATENATION",
		APIType:             f.APIType,
		DataConcatenation:   ordersMessage,
		APISchema:           "DPFMDataConcatenation",
		Accepter:            f.Accepter,
		Deleted:             f.Deleted,
		APIProcessingResult: apiErr == nil,
		APIProcessingError:  apiProcessingError,
	}
}
func (c *DPFMAPICaller) createHeaderData(f *dpfm_api_input_reader.Request, apiErr error) *dpfm_api_output_formatter.OrdersSDC {
	var apiProcessingError *string = nil
	if apiErr != nil {
		e := apiErr.Error()
		apiProcessingError = &e
	}
	return &dpfm_api_output_formatter.OrdersSDC{
		ConnectionKey:       f.ConnectionKey,
		Result:              apiErr == nil,
		RedisKey:            f.RedisKey,
		Filepath:            f.Filepath,
		APIStatusCode:       f.APIStatusCode,
		RuntimeSessionID:    f.RuntimeSessionID,
		BusinessPartnerID:   f.BusinessPartnerID,
		ServiceLabel:        "FUNCTION_ORDERS_EDI_FOR_VOLUNTARY_CHAIN_SMES_DATA_CONCATENATION",
		APIType:             f.APIType,
		DataConcatenation:   nil,
		APISchema:           "DPFMDataConcatenation",
		Accepter:            f.Accepter,
		Deleted:             f.Deleted,
		APIProcessingResult: apiErr == nil,
		APIProcessingError:  apiProcessingError,
	}
}
