package cmd

import (
	"github.com/qiniu/qshell/v2/docs"
	"github.com/qiniu/qshell/v2/iqshell/storage/object/operations"
	"github.com/spf13/cobra"
)

var batchStatCmdBuilder = func() *cobra.Command {
	var info = operations.BatchStatusInfo{}
	var cmd = &cobra.Command{
		Use:   "batchstat <Bucket> [-i <KeyListFile>]",
		Short: "Batch stat files in bucket",
		Long:  "Batch stat files in bucket, read file list from stdin if KeyListFile not specified",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				info.Bucket = args[0]
			}
			prepare(cmd, nil)
			operations.BatchStatus(info)
		},
	}
	cmd.Flags().StringVarP(&info.BatchInfo.InputFile, "input-file", "i", "", "input file")
	cmd.Flags().StringVarP(&info.BatchInfo.ItemSeparate, "sep", "F", "\t", "Separator used for split line fields")
	return cmd
}

var batchDeleteCmdBuilder = func() *cobra.Command {
	var info = operations.BatchDeleteInfo{}
	var cmd = &cobra.Command{
		Use:   "batchdelete <Bucket> [-i <KeyListFile>]",
		Short: "Batch delete files in bucket",
		Long:  "Batch delete files in bucket, read file list from stdin if KeyListFile not specified",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				info.Bucket = args[0]
			}
			prepare(cmd, nil)
			operations.BatchDelete(info)
		},
	}
	cmd.Flags().StringVarP(&info.BatchInfo.InputFile, "input-file", "i", "", "input file")
	cmd.Flags().BoolVarP(&info.BatchInfo.Force, "force", "y", false, "force mode")
	cmd.Flags().IntVarP(&info.BatchInfo.WorkCount, "worker", "c", 1, "worker count")
	cmd.Flags().StringVarP(&info.BatchInfo.SuccessExportFilePath, "success-list", "s", "", "delete success list")
	cmd.Flags().StringVarP(&info.BatchInfo.FailExportFilePath, "failure-list", "e", "", "delete failure list")
	cmd.Flags().StringVarP(&info.BatchInfo.ItemSeparate, "sep", "F", "\t", "Separator used for split line fields")
	return cmd
}

var batchChangeMimeCmdBuilder = func() *cobra.Command {
	var info = operations.BatchChangeMimeInfo{}
	var cmd = &cobra.Command{
		Use:   "batchchgm <Bucket> [-i <KeyMimeMapFile>]",
		Short: "Batch change the mime type of files in bucket",
		Long: `Batch change the mime type of files in bucket, read from stdin if KeyMimeMapFile not specified
KeyMimeMapFile content:one copy entry per line
line style:<Key><Separator><MimeType>`,
		Run: func(cmd *cobra.Command, args []string) {
			cmdId = docs.BatchChangeMimeType
			if len(args) > 0 {
				info.Bucket = args[0]
			}
			prepare(cmd, &info)
			operations.BatchChangeMime(info)
		},
	}
	cmd.Flags().StringVarP(&info.BatchInfo.InputFile, "input-file", "i", "", "input file")
	cmd.Flags().BoolVarP(&info.BatchInfo.Force, "force", "y", false, "force mode")
	cmd.Flags().IntVarP(&info.BatchInfo.WorkCount, "worker", "c", 1, "worker count")
	cmd.Flags().StringVarP(&info.BatchInfo.SuccessExportFilePath, "success-list", "s", "", "delete success list")
	cmd.Flags().StringVarP(&info.BatchInfo.FailExportFilePath, "failure-list", "e", "", "delete failure list")
	cmd.Flags().StringVarP(&info.BatchInfo.ItemSeparate, "sep", "F", "\t", "Separator used for split line fields")
	return cmd
}

var batchChangeTypeCmdBuilder = func() *cobra.Command {
	var info = operations.BatchChangeTypeInfo{}
	var cmd = &cobra.Command{
		Use:   "batchchtype <Bucket> [-i <KeyFileTypeMapFile>]",
		Short: "Batch change the file type of files in bucket",
		Long: `Batch change the file (storage) type of files in bucket, read from stdin if KeyFileTypeMapFile not specified
KeyFileTypeMapFile content:one copy entry per line
line style:<Key><Separator><type>`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				info.Bucket = args[0]
			}
			prepare(cmd, nil)
			operations.BatchChangeType(info)
		},
	}
	cmd.Flags().StringVarP(&info.BatchInfo.InputFile, "input-file", "i", "", "input file")
	cmd.Flags().BoolVarP(&info.BatchInfo.Force, "force", "y", false, "force mode")
	cmd.Flags().IntVarP(&info.BatchInfo.WorkCount, "worker", "c", 1, "worker count")
	cmd.Flags().StringVarP(&info.BatchInfo.SuccessExportFilePath, "success-list", "s", "", "delete success list")
	cmd.Flags().StringVarP(&info.BatchInfo.FailExportFilePath, "failure-list", "e", "", "delete failure list")
	cmd.Flags().StringVarP(&info.BatchInfo.ItemSeparate, "sep", "F", "\t", "Separator used for split line fields")
	return cmd
}

var batchDeleteAfterCmdBuilder = func() *cobra.Command {
	var info = operations.BatchDeleteInfo{}
	var cmd = &cobra.Command{
		Use:   "batchexpire <Bucket> [-i <KeyDeleteAfterDaysMapFile>]",
		Short: "Batch set the deleteAfterDays of the files in bucket",
		Long:  "Batch set the deleteAfterDays of the files in bucket, read from stdin if KeyDeleteAfterDaysMapFile not specified",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				info.Bucket = args[0]
			}
			prepare(cmd, nil)
			operations.BatchDeleteAfter(info)
		},
	}
	cmd.Flags().StringVarP(&info.BatchInfo.InputFile, "input-file", "i", "", "input file")
	cmd.Flags().BoolVarP(&info.BatchInfo.Force, "force", "y", false, "force mode")
	cmd.Flags().IntVarP(&info.BatchInfo.WorkCount, "worker", "c", 1, "worker count")
	cmd.Flags().StringVarP(&info.BatchInfo.ItemSeparate, "sep", "F", "\t", "Separator used for split line fields")
	return cmd
}

var batchMoveCmdBuilder = func() *cobra.Command {
	var info = operations.BatchMoveInfo{}
	var cmd = &cobra.Command{
		Use:   "batchmove <SrcBucket> <DestBucket> [-i <SrcDestKeyMapFile>]",
		Short: "Batch move files from bucket to bucket",
		Long:  "Batch move files from bucket to bucket, read from stdin if SrcDestKeyMapFile not specified",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 1 {
				info.SourceBucket = args[0]
				info.DestBucket = args[1]
			}
			prepare(cmd, nil)
			operations.BatchMove(info)
		},
	}
	cmd.Flags().StringVarP(&info.BatchInfo.InputFile, "input-file", "i", "", "input file")
	cmd.Flags().BoolVarP(&info.BatchInfo.Force, "force", "y", false, "force mode")
	cmd.Flags().BoolVarP(&info.BatchInfo.Overwrite, "overwrite", "w", false, "overwrite mode")
	cmd.Flags().IntVarP(&info.BatchInfo.WorkCount, "worker", "c", 1, "worker count")
	cmd.Flags().StringVarP(&info.BatchInfo.SuccessExportFilePath, "success-list", "s", "", "rename success list")
	cmd.Flags().StringVarP(&info.BatchInfo.FailExportFilePath, "failure-list", "e", "", "rename failure list")
	cmd.Flags().StringVarP(&info.BatchInfo.ItemSeparate, "sep", "F", "\t", "Separator used for split line fields")
	return cmd
}

var batchRenameCmdBuilder = func() *cobra.Command {
	var info = operations.BatchRenameInfo{}
	var cmd = &cobra.Command{
		Use:   "batchrename <Bucket> [-i <OldNewKeyMapFile>]",
		Short: "Batch rename files in the bucket",
		Long:  "Batch rename files in the bucket, read from stdin if OldNewKeyMapFile not specified",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				info.Bucket = args[0]
			}
			prepare(cmd, nil)
			operations.BatchRename(info)
		},
	}
	cmd.Flags().StringVarP(&info.BatchInfo.InputFile, "input-file", "i", "", "input file")
	cmd.Flags().BoolVarP(&info.BatchInfo.Force, "force", "y", false, "force mode")
	cmd.Flags().BoolVarP(&info.BatchInfo.Overwrite, "overwrite", "w", false, "overwrite mode")
	cmd.Flags().IntVarP(&info.BatchInfo.WorkCount, "worker", "c", 1, "worker count")
	cmd.Flags().StringVarP(&info.BatchInfo.SuccessExportFilePath, "success-list", "s", "", "rename success list")
	cmd.Flags().StringVarP(&info.BatchInfo.FailExportFilePath, "failure-list", "e", "", "rename failure list")
	cmd.Flags().StringVarP(&info.BatchInfo.ItemSeparate, "sep", "F", "\t", "Separator used for split line fields")
	return cmd
}

var batchCopyCmdBuilder = func() *cobra.Command {
	var info = operations.BatchCopyInfo{}
	var cmd = &cobra.Command{
		Use:   "batchcopy <SrcBucket> <DestBucket> [-i <SrcDestKeyMapFile>]",
		Short: "Batch copy files from bucket to bucket",
		Long: `Batch copy files from bucket to bucket, read from stdin if SrcDestKeyMapFile not specified.
SrcDestKeyMapFile content:one copy entry per line
line style(<ToBucketKey> use <fromBucketKey> while omitted):<fromBucketKey><Separator><ToBucketKey>
`,
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 1 {
				info.SourceBucket = args[0]
				info.DestBucket = args[1]
			}
			prepare(cmd, nil)
			operations.BatchCopy(info)
		},
	}
	cmd.Flags().StringVarP(&info.BatchInfo.InputFile, "input-file", "i", "", "input file")
	cmd.Flags().BoolVarP(&info.BatchInfo.Force, "force", "y", false, "force mode")
	cmd.Flags().BoolVarP(&info.BatchInfo.Overwrite, "overwrite", "w", false, "overwrite mode")
	cmd.Flags().IntVarP(&info.BatchInfo.WorkCount, "worker", "c", 1, "worker count")
	cmd.Flags().StringVarP(&info.BatchInfo.SuccessExportFilePath, "success-list", "s", "", "rename success list")
	cmd.Flags().StringVarP(&info.BatchInfo.FailExportFilePath, "failure-list", "e", "", "rename failure list")
	cmd.Flags().StringVarP(&info.BatchInfo.ItemSeparate, "sep", "F", "\t", "Separator used for split line fields")
	return cmd
}

var batchSignCmdBuilder = func() *cobra.Command {
	var info = operations.BatchPrivateUrlInfo{}
	var cmd = &cobra.Command{
		Use:   "batchsign [-i <ItemListFile>] [-e <Deadline>]",
		Short: "Batch create the private url from the public url list file",
		Long: `Batch create the private url from the public url list file
ItemListFile content:one copy entry per line
line style:<private url>`,
		Args: cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			prepare(cmd, nil)
			operations.BatchPrivateUrl(info)
		},
	}
	cmd.Flags().StringVarP(&info.BatchInfo.InputFile, "input-file", "i", "", "input file")
	cmd.Flags().StringVarP(&info.Deadline, "deadline", "e", "3600", "deadline in seconds")
	cmd.Flags().StringVarP(&info.BatchInfo.ItemSeparate, "sep", "F", "\t", "Separator used for split line fields")
	return cmd
}

var batchFetchCmdBuilder = func() *cobra.Command {
	var upHost = ""
	var info = operations.BatchFetchInfo{}
	var cmd = &cobra.Command{
		Use:   "batchfetch <Bucket> [-i <FetchUrlsFile>] [-c <WorkerCount>]",
		Short: "Batch fetch remoteUrls and save them in qiniu Bucket",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				info.Bucket = args[0]
			}
			if len(upHost) > 0 {
				cfg.CmdCfg.Hosts.Up = []string{upHost}
			}
			prepare(cmd, nil)
			operations.BatchFetch(info)
		},
	}
	cmd.Flags().StringVarP(&info.BatchInfo.InputFile, "input-file", "i", "", "input file")
	cmd.Flags().IntVarP(&info.BatchInfo.WorkCount, "worker", "c", 1, "worker count")
	cmd.Flags().StringVarP(&info.BatchInfo.SuccessExportFilePath, "success-list", "s", "", "rename success list")
	cmd.Flags().StringVarP(&info.BatchInfo.FailExportFilePath, "failure-list", "e", "", "rename failure list")
	cmd.Flags().StringVarP(&upHost, "up-host", "u", "", "fetch uphost")
	cmd.Flags().StringVarP(&info.BatchInfo.ItemSeparate, "sep", "F", "\t", "Separator used for split line fields")
	return cmd
}

func init() {
	rootCmd.AddCommand(
		batchStatCmdBuilder(),
		batchCopyCmdBuilder(),
		batchMoveCmdBuilder(),
		batchRenameCmdBuilder(),
		batchDeleteCmdBuilder(),
		batchDeleteAfterCmdBuilder(),
		batchChangeMimeCmdBuilder(),
		batchChangeTypeCmdBuilder(),
		batchSignCmdBuilder(),
		batchFetchCmdBuilder(),
	)
}
