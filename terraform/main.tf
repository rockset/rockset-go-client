resource "aws_s3_bucket" "rockset" {
  bucket = var.bucket
}

resource "aws_s3_bucket_object" "cities" {
  bucket = aws_s3_bucket.rockset.bucket
  key    = "cities.csv"
  content = file("${path.module}/data/cities.csv")
}

resource "aws_iam_role" "rockset" {
  name               = var.rockset-role-name
  assume_role_policy = data.aws_iam_policy_document.rockset-trust-policy.json
}

resource "aws_iam_policy" "rockset-s3-integration" {
  name = "rockset-s3-integration"
  policy = data.aws_iam_policy_document.s3-bucket-policy.json
}

data "aws_iam_policy_document" "s3-bucket-policy" {
  statement {
    sid       = "RocksetS3IntegrationPolicy"
    actions   = [
      "s3:List*",
      "s3:GetObject"
    ]
    resources = [
      aws_s3_bucket.rockset.arn,
      "${aws_s3_bucket.rockset.arn}/*"
    ]
  }
}

data "aws_iam_policy_document" "rockset-trust-policy" {
  statement {
    sid     = ""
    effect  = "Allow"
    actions = [
      "sts:AssumeRole"]
    principals {
      identifiers = [
        "arn:aws:iam::${var.rockset-account-id}:root"]
      type        = "AWS"
    }
    condition {
      test     = "StringEquals"
      values   = [
        var.rockset-external-id]
      variable = "sts:ExternalId"
    }
  }
}

resource "aws_iam_role_policy_attachment" "rockset-s3-integration" {
  role       = aws_iam_role.rockset.name
  policy_arn = aws_iam_policy.rockset-s3-integration.arn
}