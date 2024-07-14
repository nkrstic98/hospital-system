package constants

import "hospital-system/server/app/dto"

const (
	TestType_CBC          = "Complete Blood Count (CBC)"
	TestType_BMP          = "Basic Metabolic Panel (BMP)"
	TestType_CMP          = "Comprehensive Metabolic Panel (CMP)"
	TestType_LipidPanel   = "Lipid Panel"
	TestType_LFT          = "Liver Function Tests (LFTs)"
	TestType_Thyroid      = "Thyroid Function Tests"
	TestType_Coagulation  = "Coagulation Tests"
	TestType_Urinalysis   = "Urinalysis"
	TestType_Hemoglobin   = "Hemoglobin A1c (HbA1c)"
	TestType_BloodCulture = "Blood Culture"
	TestType_CRP          = "C-Reactive Protein (CRP)"
	TestType_ESR          = "Erythrocyte Sedimentation Rate (ESR)"
	TestType_Vitamins     = "Vitamin and Mineral Tests"
	TestType_Hormones     = "Hormone Tests"
	TestType_IDT          = "Infectious Disease Tests"
	TestType_TumorMarkers = "Tumor Markers"
)

var LabTests = map[string][]dto.LabTest{
	"Complete Blood Count (CBC)": {
		{Name: "White Blood Cells (WBC)", Unit: "cells/mcL", MinValue: 4500, MaxValue: 11000, ReferenceRange: "4500-11000 cells/mcL"},
		{Name: "Red Blood Cells (RBC)", Unit: "cells/mcL", MinValue: 4.5, MaxValue: 6.0, ReferenceRange: "4.5-6.0 cells/mcL"},
		{Name: "Hemoglobin (Hgb)", Unit: "g/dL", MinValue: 13.5, MaxValue: 17.5, ReferenceRange: "13.5-17.5 g/dL"},
		{Name: "Hematocrit (Hct)", Unit: "%", MinValue: 40, MaxValue: 50, ReferenceRange: "40-50 %"},
		{Name: "Platelets", Unit: "cells/mcL", MinValue: 150000, MaxValue: 450000, ReferenceRange: "150000-450000 cells/mcL"},
	},
	"Basic Metabolic Panel (BMP)": {
		{Name: "Sodium", Unit: "mEq/L", MinValue: 135, MaxValue: 145, ReferenceRange: "135-145 mEq/L"},
		{Name: "Potassium", Unit: "mEq/L", MinValue: 3.5, MaxValue: 5.0, ReferenceRange: "3.5-5.0 mEq/L"},
		{Name: "Chloride", Unit: "mEq/L", MinValue: 96, MaxValue: 106, ReferenceRange: "96-106 mEq/L"},
		{Name: "Bicarbonate", Unit: "mEq/L", MinValue: 22, MaxValue: 28, ReferenceRange: "22-28 mEq/L"},
		{Name: "Blood Urea Nitrogen (BUN)", Unit: "mg/dL", MinValue: 6, MaxValue: 20, ReferenceRange: "6-20 mg/dL"},
		{Name: "Creatinine", Unit: "mg/dL", MinValue: 0.6, MaxValue: 1.2, ReferenceRange: "0.6-1.2 mg/dL"},
		{Name: "Glucose", Unit: "mg/dL", MinValue: 70, MaxValue: 100, ReferenceRange: "70-100 mg/dL"},
		{Name: "Calcium", Unit: "mg/dL", MinValue: 8.5, MaxValue: 10.5, ReferenceRange: "8.5-10.5 mg/dL"},
	},
	"Comprehensive Metabolic Panel (CMP)": {
		{Name: "Sodium", Unit: "mEq/L", MinValue: 135, MaxValue: 145, ReferenceRange: "135-145 mEq/L"},
		{Name: "Potassium", Unit: "mEq/L", MinValue: 3.5, MaxValue: 5.0, ReferenceRange: "3.5-5.0 mEq/L"},
		{Name: "Chloride", Unit: "mEq/L", MinValue: 96, MaxValue: 106, ReferenceRange: "96-106 mEq/L"},
		{Name: "Bicarbonate", Unit: "mEq/L", MinValue: 22, MaxValue: 28, ReferenceRange: "22-28 mEq/L"},
		{Name: "Blood Urea Nitrogen (BUN)", Unit: "mg/dL", MinValue: 6, MaxValue: 20, ReferenceRange: "6-20 mg/dL"},
		{Name: "Creatinine", Unit: "mg/dL", MinValue: 0.6, MaxValue: 1.2, ReferenceRange: "0.6-1.2 mg/dL"},
		{Name: "Glucose", Unit: "mg/dL", MinValue: 70, MaxValue: 100, ReferenceRange: "70-100 mg/dL"},
		{Name: "Calcium", Unit: "mg/dL", MinValue: 8.5, MaxValue: 10.5, ReferenceRange: "8.5-10.5 mg/dL"},
		{Name: "Total Protein", Unit: "g/dL", MinValue: 6.0, MaxValue: 8.3, ReferenceRange: "6.0-8.3 g/dL"},
		{Name: "Albumin", Unit: "g/dL", MinValue: 3.5, MaxValue: 5.0, ReferenceRange: "3.5-5.0 g/dL"},
		{Name: "Bilirubin", Unit: "mg/dL", MinValue: 0.1, MaxValue: 1.2, ReferenceRange: "0.1-1.2 mg/dL"},
		{Name: "Alkaline Phosphatase (ALP)", Unit: "U/L", MinValue: 44, MaxValue: 147, ReferenceRange: "44-147 U/L"},
		{Name: "Aspartate Aminotransferase (AST)", Unit: "U/L", MinValue: 10, MaxValue: 40, ReferenceRange: "10-40 U/L"},
		{Name: "Alanine Aminotransferase (ALT)", Unit: "U/L", MinValue: 7, MaxValue: 56, ReferenceRange: "7-56 U/L"},
	},
	"Lipid Panel": {
		{Name: "Total Cholesterol", Unit: "mg/dL", MinValue: 125, MaxValue: 200, ReferenceRange: "125-200 mg/dL"},
		{Name: "High-Density Lipoprotein (HDL)", Unit: "mg/dL", MinValue: 40, MaxValue: 60, ReferenceRange: "40-60 mg/dL"},
		{Name: "Low-Density Lipoprotein (LDL)", Unit: "mg/dL", MinValue: 0, MaxValue: 100, ReferenceRange: "0-100 mg/dL"},
		{Name: "Triglycerides", Unit: "mg/dL", MinValue: 0, MaxValue: 150, ReferenceRange: "0-150 mg/dL"},
	},
	"Liver Function Tests (LFTs)": {
		{Name: "Total Protein", Unit: "g/dL", MinValue: 6.0, MaxValue: 8.3, ReferenceRange: "6.0-8.3 g/dL"},
		{Name: "Albumin", Unit: "g/dL", MinValue: 3.5, MaxValue: 5.0, ReferenceRange: "3.5-5.0 g/dL"},
		{Name: "Bilirubin", Unit: "mg/dL", MinValue: 0.1, MaxValue: 1.2, ReferenceRange: "0.1-1.2 mg/dL"},
		{Name: "Alkaline Phosphatase (ALP)", Unit: "U/L", MinValue: 44, MaxValue: 147, ReferenceRange: "44-147 U/L"},
		{Name: "Aspartate Aminotransferase (AST)", Unit: "U/L", MinValue: 10, MaxValue: 40, ReferenceRange: "10-40 U/L"},
		{Name: "Alanine Aminotransferase (ALT)", Unit: "U/L", MinValue: 7, MaxValue: 56, ReferenceRange: "7-56 U/L"},
	},
	"Thyroid Function Tests": {
		{Name: "Thyroid Stimulating Hormone (TSH)", Unit: "mIU/L", MinValue: 0.4, MaxValue: 4.0, ReferenceRange: "0.4-4.0 mIU/L"},
		{Name: "Triiodothyronine (T3)", Unit: "ng/dL", MinValue: 100, MaxValue: 200, ReferenceRange: "100-200 ng/dL"},
		{Name: "Thyroxine (T4)", Unit: "ng/dL", MinValue: 5.0, MaxValue: 12.0, ReferenceRange: "5.0-12.0 ng/dL"},
	},
	"Coagulation Tests": {
		{Name: "Prothrombin Time (PT)", Unit: "seconds", MinValue: 11, MaxValue: 13.5, ReferenceRange: "11-13.5 seconds"},
		{Name: "International Normalized Ratio (INR)", Unit: "", MinValue: 0.8, MaxValue: 1.2, ReferenceRange: "0.8-1.2"},
		{Name: "Activated Partial Thromboplastin Time (aPTT)", Unit: "seconds", MinValue: 25, MaxValue: 35, ReferenceRange: "25-35 seconds"},
	},
	"Urinalysis": {
		{Name: "Specific Gravity", Unit: "", MinValue: 1.005, MaxValue: 1.030, ReferenceRange: "1.005-1.030"},
		{Name: "pH", Unit: "", MinValue: 4.6, MaxValue: 8.0, ReferenceRange: "4.6-8.0"},
		{Name: "Protein", Unit: "mg/dL", MinValue: 0, MaxValue: 8, ReferenceRange: "0-8 mg/dL"},
		{Name: "Glucose", Unit: "mg/dL", MinValue: 0, MaxValue: 15, ReferenceRange: "0-15 mg/dL"},
		{Name: "Ketones", Unit: "mg/dL", MinValue: 0, MaxValue: 3, ReferenceRange: "0-3 mg/dL"},
	},
	"Hemoglobin A1c (HbA1c)": {
		{Name: "Hemoglobin A1c (HbA1c)", Unit: "%", MinValue: 4, MaxValue: 5.6, ReferenceRange: "4-5.6 %"},
	},
	"Blood Culture": {
		{Name: "Blood Culture", Unit: "positive/negative", MinValue: 0, MaxValue: 1, ReferenceRange: "positive/negative"},
	},
	"C-Reactive Protein (CRP)": {
		{Name: "C-Reactive Protein (CRP)", Unit: "mg/L", MinValue: 0, MaxValue: 10, ReferenceRange: "0-10 mg/L"},
	},
	"Erythrocyte Sedimentation Rate (ESR)": {
		{Name: "Erythrocyte Sedimentation Rate (ESR)", Unit: "mm/hr", MinValue: 0, MaxValue: 20, ReferenceRange: "0-20 mm/hr"},
	},
	"Vitamin and Mineral Tests": {
		{Name: "Vitamin D", Unit: "ng/mL", MinValue: 20, MaxValue: 50, ReferenceRange: "20-50 ng/mL"},
		{Name: "Vitamin B12", Unit: "pg/mL", MinValue: 200, MaxValue: 900, ReferenceRange: "200-900 pg/mL"},
		{Name: "Iron", Unit: "mcg/dL", MinValue: 60, MaxValue: 170, ReferenceRange: "60-170 mcg/dL"},
	},
	"Hormone Tests": {
		{Name: "Cortisol", Unit: "mcg/dL", MinValue: 6, MaxValue: 23, ReferenceRange: "6-23 mcg/dL"},
		{Name: "Estrogen", Unit: "pg/mL", MinValue: 15, MaxValue: 350, ReferenceRange: "15-350 pg/mL"},
		{Name: "Testosterone", Unit: "ng/dL", MinValue: 300, MaxValue: 1000, ReferenceRange: "300-1000 ng/dL"},
	},
	"Infectious Disease Tests": {
		{Name: "HIV", Unit: "positive/negative", MinValue: 0, MaxValue: 1, ReferenceRange: "positive/negative"},
		{Name: "Hepatitis B", Unit: "positive/negative", MinValue: 0, MaxValue: 1, ReferenceRange: "positive/negative"},
		{Name: "Chlamydia", Unit: "positive/negative", MinValue: 0, MaxValue: 1, ReferenceRange: "positive/negative"},
		{Name: "Gonorrhea", Unit: "positive/negative", MinValue: 0, MaxValue: 1, ReferenceRange: "positive/negative"},
		{Name: "Syphilis", Unit: "positive/negative", MinValue: 0, MaxValue: 1, ReferenceRange: "positive/negative"},
		{Name: "Herpes", Unit: "positive/negative", MinValue: 0, MaxValue: 1, ReferenceRange: "positive/negative"},
		{Name: "HPV", Unit: "positive/negative", MinValue: 0, MaxValue: 1, ReferenceRange: "positive/negative"},
		{Name: "Trichomoniasis", Unit: "positive/negative", MinValue: 0, MaxValue: 1, ReferenceRange: "positive/negative"},
	},
	"Tumor Markers": {
		{Name: "Prostate-Specific Antigen (PSA)", Unit: "ng/mL", MinValue: 0, MaxValue: 4, ReferenceRange: "0-4 ng/mL"},
		{Name: "Cancer Antigen 125 (CA-125)", Unit: "U/mL", MinValue: 0, MaxValue: 35, ReferenceRange: "0-35 U/mL"},
		{Name: "Carcinoembryonic Antigen (CEA)", Unit: "ng/mL", MinValue: 0, MaxValue: 5, ReferenceRange: "0-5 ng/mL"},
	},
}
