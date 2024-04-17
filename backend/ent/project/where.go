// Code generated by ent, DO NOT EDIT.

package project

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/cble-platform/cble/backend/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldName, v))
}

// QuotaCPU applies equality check predicate on the "quota_cpu" field. It's identical to QuotaCPUEQ.
func QuotaCPU(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldQuotaCPU, v))
}

// UsageCPU applies equality check predicate on the "usage_cpu" field. It's identical to UsageCPUEQ.
func UsageCPU(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUsageCPU, v))
}

// QuotaRAM applies equality check predicate on the "quota_ram" field. It's identical to QuotaRAMEQ.
func QuotaRAM(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldQuotaRAM, v))
}

// UsageRAM applies equality check predicate on the "usage_ram" field. It's identical to UsageRAMEQ.
func UsageRAM(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUsageRAM, v))
}

// QuotaDisk applies equality check predicate on the "quota_disk" field. It's identical to QuotaDiskEQ.
func QuotaDisk(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldQuotaDisk, v))
}

// UsageDisk applies equality check predicate on the "usage_disk" field. It's identical to UsageDiskEQ.
func UsageDisk(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUsageDisk, v))
}

// QuotaNetwork applies equality check predicate on the "quota_network" field. It's identical to QuotaNetworkEQ.
func QuotaNetwork(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldQuotaNetwork, v))
}

// UsageNetwork applies equality check predicate on the "usage_network" field. It's identical to UsageNetworkEQ.
func UsageNetwork(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUsageNetwork, v))
}

// QuotaRouter applies equality check predicate on the "quota_router" field. It's identical to QuotaRouterEQ.
func QuotaRouter(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldQuotaRouter, v))
}

// UsageRouter applies equality check predicate on the "usage_router" field. It's identical to UsageRouterEQ.
func UsageRouter(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUsageRouter, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldName, v))
}

// QuotaCPUEQ applies the EQ predicate on the "quota_cpu" field.
func QuotaCPUEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldQuotaCPU, v))
}

// QuotaCPUNEQ applies the NEQ predicate on the "quota_cpu" field.
func QuotaCPUNEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldQuotaCPU, v))
}

// QuotaCPUIn applies the In predicate on the "quota_cpu" field.
func QuotaCPUIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldQuotaCPU, vs...))
}

// QuotaCPUNotIn applies the NotIn predicate on the "quota_cpu" field.
func QuotaCPUNotIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldQuotaCPU, vs...))
}

// QuotaCPUGT applies the GT predicate on the "quota_cpu" field.
func QuotaCPUGT(v int) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldQuotaCPU, v))
}

// QuotaCPUGTE applies the GTE predicate on the "quota_cpu" field.
func QuotaCPUGTE(v int) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldQuotaCPU, v))
}

// QuotaCPULT applies the LT predicate on the "quota_cpu" field.
func QuotaCPULT(v int) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldQuotaCPU, v))
}

// QuotaCPULTE applies the LTE predicate on the "quota_cpu" field.
func QuotaCPULTE(v int) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldQuotaCPU, v))
}

// UsageCPUEQ applies the EQ predicate on the "usage_cpu" field.
func UsageCPUEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUsageCPU, v))
}

// UsageCPUNEQ applies the NEQ predicate on the "usage_cpu" field.
func UsageCPUNEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldUsageCPU, v))
}

// UsageCPUIn applies the In predicate on the "usage_cpu" field.
func UsageCPUIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldUsageCPU, vs...))
}

// UsageCPUNotIn applies the NotIn predicate on the "usage_cpu" field.
func UsageCPUNotIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldUsageCPU, vs...))
}

// UsageCPUGT applies the GT predicate on the "usage_cpu" field.
func UsageCPUGT(v int) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldUsageCPU, v))
}

// UsageCPUGTE applies the GTE predicate on the "usage_cpu" field.
func UsageCPUGTE(v int) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldUsageCPU, v))
}

// UsageCPULT applies the LT predicate on the "usage_cpu" field.
func UsageCPULT(v int) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldUsageCPU, v))
}

// UsageCPULTE applies the LTE predicate on the "usage_cpu" field.
func UsageCPULTE(v int) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldUsageCPU, v))
}

// QuotaRAMEQ applies the EQ predicate on the "quota_ram" field.
func QuotaRAMEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldQuotaRAM, v))
}

// QuotaRAMNEQ applies the NEQ predicate on the "quota_ram" field.
func QuotaRAMNEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldQuotaRAM, v))
}

// QuotaRAMIn applies the In predicate on the "quota_ram" field.
func QuotaRAMIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldQuotaRAM, vs...))
}

// QuotaRAMNotIn applies the NotIn predicate on the "quota_ram" field.
func QuotaRAMNotIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldQuotaRAM, vs...))
}

// QuotaRAMGT applies the GT predicate on the "quota_ram" field.
func QuotaRAMGT(v int) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldQuotaRAM, v))
}

// QuotaRAMGTE applies the GTE predicate on the "quota_ram" field.
func QuotaRAMGTE(v int) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldQuotaRAM, v))
}

// QuotaRAMLT applies the LT predicate on the "quota_ram" field.
func QuotaRAMLT(v int) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldQuotaRAM, v))
}

// QuotaRAMLTE applies the LTE predicate on the "quota_ram" field.
func QuotaRAMLTE(v int) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldQuotaRAM, v))
}

// UsageRAMEQ applies the EQ predicate on the "usage_ram" field.
func UsageRAMEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUsageRAM, v))
}

// UsageRAMNEQ applies the NEQ predicate on the "usage_ram" field.
func UsageRAMNEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldUsageRAM, v))
}

// UsageRAMIn applies the In predicate on the "usage_ram" field.
func UsageRAMIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldUsageRAM, vs...))
}

// UsageRAMNotIn applies the NotIn predicate on the "usage_ram" field.
func UsageRAMNotIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldUsageRAM, vs...))
}

// UsageRAMGT applies the GT predicate on the "usage_ram" field.
func UsageRAMGT(v int) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldUsageRAM, v))
}

// UsageRAMGTE applies the GTE predicate on the "usage_ram" field.
func UsageRAMGTE(v int) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldUsageRAM, v))
}

// UsageRAMLT applies the LT predicate on the "usage_ram" field.
func UsageRAMLT(v int) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldUsageRAM, v))
}

// UsageRAMLTE applies the LTE predicate on the "usage_ram" field.
func UsageRAMLTE(v int) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldUsageRAM, v))
}

// QuotaDiskEQ applies the EQ predicate on the "quota_disk" field.
func QuotaDiskEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldQuotaDisk, v))
}

// QuotaDiskNEQ applies the NEQ predicate on the "quota_disk" field.
func QuotaDiskNEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldQuotaDisk, v))
}

// QuotaDiskIn applies the In predicate on the "quota_disk" field.
func QuotaDiskIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldQuotaDisk, vs...))
}

// QuotaDiskNotIn applies the NotIn predicate on the "quota_disk" field.
func QuotaDiskNotIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldQuotaDisk, vs...))
}

// QuotaDiskGT applies the GT predicate on the "quota_disk" field.
func QuotaDiskGT(v int) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldQuotaDisk, v))
}

// QuotaDiskGTE applies the GTE predicate on the "quota_disk" field.
func QuotaDiskGTE(v int) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldQuotaDisk, v))
}

// QuotaDiskLT applies the LT predicate on the "quota_disk" field.
func QuotaDiskLT(v int) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldQuotaDisk, v))
}

// QuotaDiskLTE applies the LTE predicate on the "quota_disk" field.
func QuotaDiskLTE(v int) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldQuotaDisk, v))
}

// UsageDiskEQ applies the EQ predicate on the "usage_disk" field.
func UsageDiskEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUsageDisk, v))
}

// UsageDiskNEQ applies the NEQ predicate on the "usage_disk" field.
func UsageDiskNEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldUsageDisk, v))
}

// UsageDiskIn applies the In predicate on the "usage_disk" field.
func UsageDiskIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldUsageDisk, vs...))
}

// UsageDiskNotIn applies the NotIn predicate on the "usage_disk" field.
func UsageDiskNotIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldUsageDisk, vs...))
}

// UsageDiskGT applies the GT predicate on the "usage_disk" field.
func UsageDiskGT(v int) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldUsageDisk, v))
}

// UsageDiskGTE applies the GTE predicate on the "usage_disk" field.
func UsageDiskGTE(v int) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldUsageDisk, v))
}

// UsageDiskLT applies the LT predicate on the "usage_disk" field.
func UsageDiskLT(v int) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldUsageDisk, v))
}

// UsageDiskLTE applies the LTE predicate on the "usage_disk" field.
func UsageDiskLTE(v int) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldUsageDisk, v))
}

// QuotaNetworkEQ applies the EQ predicate on the "quota_network" field.
func QuotaNetworkEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldQuotaNetwork, v))
}

// QuotaNetworkNEQ applies the NEQ predicate on the "quota_network" field.
func QuotaNetworkNEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldQuotaNetwork, v))
}

// QuotaNetworkIn applies the In predicate on the "quota_network" field.
func QuotaNetworkIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldQuotaNetwork, vs...))
}

// QuotaNetworkNotIn applies the NotIn predicate on the "quota_network" field.
func QuotaNetworkNotIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldQuotaNetwork, vs...))
}

// QuotaNetworkGT applies the GT predicate on the "quota_network" field.
func QuotaNetworkGT(v int) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldQuotaNetwork, v))
}

// QuotaNetworkGTE applies the GTE predicate on the "quota_network" field.
func QuotaNetworkGTE(v int) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldQuotaNetwork, v))
}

// QuotaNetworkLT applies the LT predicate on the "quota_network" field.
func QuotaNetworkLT(v int) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldQuotaNetwork, v))
}

// QuotaNetworkLTE applies the LTE predicate on the "quota_network" field.
func QuotaNetworkLTE(v int) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldQuotaNetwork, v))
}

// UsageNetworkEQ applies the EQ predicate on the "usage_network" field.
func UsageNetworkEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUsageNetwork, v))
}

// UsageNetworkNEQ applies the NEQ predicate on the "usage_network" field.
func UsageNetworkNEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldUsageNetwork, v))
}

// UsageNetworkIn applies the In predicate on the "usage_network" field.
func UsageNetworkIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldUsageNetwork, vs...))
}

// UsageNetworkNotIn applies the NotIn predicate on the "usage_network" field.
func UsageNetworkNotIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldUsageNetwork, vs...))
}

// UsageNetworkGT applies the GT predicate on the "usage_network" field.
func UsageNetworkGT(v int) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldUsageNetwork, v))
}

// UsageNetworkGTE applies the GTE predicate on the "usage_network" field.
func UsageNetworkGTE(v int) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldUsageNetwork, v))
}

// UsageNetworkLT applies the LT predicate on the "usage_network" field.
func UsageNetworkLT(v int) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldUsageNetwork, v))
}

// UsageNetworkLTE applies the LTE predicate on the "usage_network" field.
func UsageNetworkLTE(v int) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldUsageNetwork, v))
}

// QuotaRouterEQ applies the EQ predicate on the "quota_router" field.
func QuotaRouterEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldQuotaRouter, v))
}

// QuotaRouterNEQ applies the NEQ predicate on the "quota_router" field.
func QuotaRouterNEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldQuotaRouter, v))
}

// QuotaRouterIn applies the In predicate on the "quota_router" field.
func QuotaRouterIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldQuotaRouter, vs...))
}

// QuotaRouterNotIn applies the NotIn predicate on the "quota_router" field.
func QuotaRouterNotIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldQuotaRouter, vs...))
}

// QuotaRouterGT applies the GT predicate on the "quota_router" field.
func QuotaRouterGT(v int) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldQuotaRouter, v))
}

// QuotaRouterGTE applies the GTE predicate on the "quota_router" field.
func QuotaRouterGTE(v int) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldQuotaRouter, v))
}

// QuotaRouterLT applies the LT predicate on the "quota_router" field.
func QuotaRouterLT(v int) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldQuotaRouter, v))
}

// QuotaRouterLTE applies the LTE predicate on the "quota_router" field.
func QuotaRouterLTE(v int) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldQuotaRouter, v))
}

// UsageRouterEQ applies the EQ predicate on the "usage_router" field.
func UsageRouterEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUsageRouter, v))
}

// UsageRouterNEQ applies the NEQ predicate on the "usage_router" field.
func UsageRouterNEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldUsageRouter, v))
}

// UsageRouterIn applies the In predicate on the "usage_router" field.
func UsageRouterIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldUsageRouter, vs...))
}

// UsageRouterNotIn applies the NotIn predicate on the "usage_router" field.
func UsageRouterNotIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldUsageRouter, vs...))
}

// UsageRouterGT applies the GT predicate on the "usage_router" field.
func UsageRouterGT(v int) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldUsageRouter, v))
}

// UsageRouterGTE applies the GTE predicate on the "usage_router" field.
func UsageRouterGTE(v int) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldUsageRouter, v))
}

// UsageRouterLT applies the LT predicate on the "usage_router" field.
func UsageRouterLT(v int) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldUsageRouter, v))
}

// UsageRouterLTE applies the LTE predicate on the "usage_router" field.
func UsageRouterLTE(v int) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldUsageRouter, v))
}

// HasMembers applies the HasEdge predicate on the "members" edge.
func HasMembers() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, MembersTable, MembersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMembersWith applies the HasEdge predicate on the "members" edge with a given conditions (other predicates).
func HasMembersWith(preds ...predicate.User) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newMembersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGroupMembers applies the HasEdge predicate on the "group_members" edge.
func HasGroupMembers() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, GroupMembersTable, GroupMembersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupMembersWith applies the HasEdge predicate on the "group_members" edge with a given conditions (other predicates).
func HasGroupMembersWith(preds ...predicate.Group) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newGroupMembersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBlueprints applies the HasEdge predicate on the "blueprints" edge.
func HasBlueprints() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, BlueprintsTable, BlueprintsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBlueprintsWith applies the HasEdge predicate on the "blueprints" edge with a given conditions (other predicates).
func HasBlueprintsWith(preds ...predicate.Blueprint) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newBlueprintsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDeployments applies the HasEdge predicate on the "deployments" edge.
func HasDeployments() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, DeploymentsTable, DeploymentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeploymentsWith applies the HasEdge predicate on the "deployments" edge with a given conditions (other predicates).
func HasDeploymentsWith(preds ...predicate.Deployment) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newDeploymentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMemberships applies the HasEdge predicate on the "memberships" edge.
func HasMemberships() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, MembershipsTable, MembershipsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMembershipsWith applies the HasEdge predicate on the "memberships" edge with a given conditions (other predicates).
func HasMembershipsWith(preds ...predicate.Membership) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newMembershipsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGroupMemberships applies the HasEdge predicate on the "group_memberships" edge.
func HasGroupMemberships() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, GroupMembershipsTable, GroupMembershipsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupMembershipsWith applies the HasEdge predicate on the "group_memberships" edge with a given conditions (other predicates).
func HasGroupMembershipsWith(preds ...predicate.GroupMembership) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newGroupMembershipsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Project) predicate.Project {
	return predicate.Project(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Project) predicate.Project {
	return predicate.Project(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Project) predicate.Project {
	return predicate.Project(sql.NotPredicates(p))
}
