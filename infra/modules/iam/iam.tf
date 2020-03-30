resource "aws_iam_role" "lambda_role" {
  name = "${var.app_name}-lambda-role"

  assume_role_policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow"
    }
  ]
}
POLICY
}

# Lambda Policy Data
data "aws_iam_policy_document" "lambda_policy_document" {
  statement {
    actions = [
      "s3:PutObject",

      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]

    effect    = "Allow"
    resources = ["*"]
  }
}

# Lambda Attach Role to Policy
resource "aws_iam_role_policy" "lambda_role_policy" {
  role   = aws_iam_role.lambda_role.id
  name   = "${var.app_name}-lambda-policy"
  policy = data.aws_iam_policy_document.lambda_policy_document.json
}
