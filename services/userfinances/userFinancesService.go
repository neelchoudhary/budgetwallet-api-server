package userfinances

import (
	context "context"
	fmt "fmt"

	"github.com/neelchoudhary/budgetmanagergrpc/models"
	shared "github.com/neelchoudhary/budgetmanagergrpc/services/shared"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Service UserFinancesService struct
type Service struct {
	financialItemRepo    models.FinancialItemRepository
	financialAccountRepo models.FinancialAccountRepository
}

// NewUserFinancesServer contructor to assign repo
func NewUserFinancesServer(itemRepo *models.FinancialItemRepository, accountRepo *models.FinancialAccountRepository) UserFinancesServiceServer {
	return &Service{financialItemRepo: *itemRepo, financialAccountRepo: *accountRepo}
}

// GetFinancialInstitutions get financial institutions from DB for a user
func (s *Service) GetFinancialInstitutions(ctx context.Context, req *GetFinancialInstitutionsRequest) (*GetFinancialInstitutionsResponse, error) {
	items, err := s.financialItemRepo.GetUserItems(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Repo error getting items: %s", err.Error()))
	}

	var pbItems []*shared.FinancialInstitution
	for _, item := range items {
		pbItems = append(pbItems, dataToItemPb(item))
	}

	res := &GetFinancialInstitutionsResponse{
		FinancialInstitutions: pbItems,
	}
	return res, nil
}

// GetFinancialAccounts get financial accounts from DB for a user's item
func (s *Service) GetFinancialAccounts(ctx context.Context, req *GetFinancialAccountsRequest) (*GetFinancialAccountsResponse, error) {
	accounts, err := s.financialAccountRepo.GetItemAccounts(req.UserId, req.ItemId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Repo error getting accounts: %s", err.Error()))
	}
	var pbAccounts []*shared.FinancialAccount
	for _, account := range accounts {
		pbAccounts = append(pbAccounts, dataToAccountPb(account))
	}

	res := &GetFinancialAccountsResponse{
		FinancialAccounts: pbAccounts,
	}
	return res, nil
}

// ToggleFinancialAccount toggle the selected property of a financial account for a user's item
func (s *Service) ToggleFinancialAccount(ctx context.Context, req *ToggleFinancialAccountRequest) (*ToggleFinancialAccountResponse, error) {
	account, err := s.financialAccountRepo.GetAccountByID(req.GetUserId(), req.GetAccountId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Repo error getting account by id: %s", err.Error()))
	}
	account.SetSelected(req.GetSelected())
	err = s.financialAccountRepo.UpdateAccount(req.GetUserId(), req.GetAccountId(), account)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Repo error updatinig account: %s", err.Error()))
	}
	res := &ToggleFinancialAccountResponse{
		Success: true,
	}
	return res, nil
}

// GetTransactions get all transactions for a user's item
func (s *Service) GetTransactions(ctx context.Context, req *GetTransactionsRequest) (*GetTransactionsResponse, error) {
	return nil, nil
}

func dataToItemPb(data models.FinancialItem) *shared.FinancialInstitution {
	return &shared.FinancialInstitution{
		Id:               data.ID,
		InstitutionName:  data.InstitutionName,
		InstitutionColor: data.InstitutionColor,
		InstitutionLogo:  data.InstitutionLogo,
	}
}

func dataToAccountPb(data models.FinancialAccount) *shared.FinancialAccount {
	return &shared.FinancialAccount{
		Id:             data.ID,
		UserId:         data.UserID,
		ItemId:         data.ItemID,
		PlaidAccountId: data.PlaidAccountID,
		CurrentBalance: data.CurrentBalance,
		AccountName:    data.AccountName,
		OfficialName:   data.OfficialName,
		AccountType:    data.AccountType,
		AccountSubtype: data.AccountSubType,
		AccountMask:    data.AccountMask,
		Selected:       data.Selected,
	}
}