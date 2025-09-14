// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package avcodec contains the codecs (decoders and encoders) provided by the libavcodec library
// Provides some generic global options, which can be set on all the encoders and decoders.
package avcodec

//#cgo pkg-config: libavformat libavcodec libavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
//#include <libswscale/swscale.h>
import "C"

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage
const (
	AV_PIX_FMT_NONE      = C.AV_PIX_FMT_NONE
	AV_PIX_FMT_YUV420P   = C.AV_PIX_FMT_YUV420P   ///< planar YUV 4:2:0, 12bpp, (1 Cr & Cb sample per 2x2 Y samples)
	AV_PIX_FMT_YUYV422   = C.AV_PIX_FMT_YUYV422   ///< packed YUV 4:2:2, 16bpp, Y0 Cb Y1 Cr
	AV_PIX_FMT_RGB24     = C.AV_PIX_FMT_RGB24     ///< packed RGB 8:8:8, 24bpp, RGBRGB...
	AV_PIX_FMT_BGR24     = C.AV_PIX_FMT_BGR24     ///< packed RGB 8:8:8, 24bpp, BGRBGR...
	AV_PIX_FMT_YUV422P   = C.AV_PIX_FMT_YUV422P   ///< planar YUV 4:2:2, 16bpp, (1 Cr & Cb sample per 2x1 Y samples)
	AV_PIX_FMT_YUV444P   = C.AV_PIX_FMT_YUV444P   ///< planar YUV 4:4:4, 24bpp, (1 Cr & Cb sample per 1x1 Y samples)
	AV_PIX_FMT_YUV410P   = C.AV_PIX_FMT_YUV410P   ///< planar YUV 4:1:0,  9bpp, (1 Cr & Cb sample per 4x4 Y samples)
	AV_PIX_FMT_YUV411P   = C.AV_PIX_FMT_YUV411P   ///< planar YUV 4:1:1, 12bpp, (1 Cr & Cb sample per 4x1 Y samples)
	AV_PIX_FMT_GRAY8     = C.AV_PIX_FMT_GRAY8     ///<        Y        ,  8bpp
	AV_PIX_FMT_MONOWHITE = C.AV_PIX_FMT_MONOWHITE ///<        Y        ,  1bpp, 0 is white, 1 is black, in each byte pixels are ordered from the msb to the lsb
	AV_PIX_FMT_MONOBLACK = C.AV_PIX_FMT_MONOBLACK ///<        Y        ,  1bpp, 0 is black, 1 is white, in each byte pixels are ordered from the msb to the lsb
	AV_PIX_FMT_PAL8      = C.AV_PIX_FMT_PAL8      ///< 8 bits with AV_PIX_FMT_RGB32 palette
	AV_PIX_FMT_YUVJ420P  = C.AV_PIX_FMT_YUVJ420P  ///< planar YUV 4:2:0, 12bpp, full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV420P and setting color_range
	AV_PIX_FMT_YUVJ422P  = C.AV_PIX_FMT_YUVJ422P  ///< planar YUV 4:2:2, 16bpp, full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV422P and setting color_range
	AV_PIX_FMT_YUVJ444P  = C.AV_PIX_FMT_YUVJ444P  ///< planar YUV 4:4:4, 24bpp, full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV444P and setting color_range
	AV_PIX_FMT_UYVY422   = C.AV_PIX_FMT_UYVY422   ///< packed YUV 4:2:2, 16bpp, Cb Y0 Cr Y1
	AV_PIX_FMT_UYYVYY411 = C.AV_PIX_FMT_UYYVYY411 ///< packed YUV 4:1:1, 12bpp, Cb Y0 Y1 Cr Y2 Y3
	AV_PIX_FMT_BGR8      = C.AV_PIX_FMT_BGR8      ///< packed RGB 3:3:2,  8bpp, (msb)2B 3G 3R(lsb)
	AV_PIX_FMT_BGR4      = C.AV_PIX_FMT_BGR4      ///< packed RGB 1:2:1 bitstream,  4bpp, (msb)1B 2G 1R(lsb), a byte contains two pixels, the first pixel in the byte is the one composed by the 4 msb bits
	AV_PIX_FMT_BGR4_BYTE = C.AV_PIX_FMT_BGR4_BYTE ///< packed RGB 1:2:1,  8bpp, (msb)1B 2G 1R(lsb)
	AV_PIX_FMT_RGB8      = C.AV_PIX_FMT_RGB8      ///< packed RGB 3:3:2,  8bpp, (msb)2R 3G 3B(lsb)
	AV_PIX_FMT_RGB4      = C.AV_PIX_FMT_RGB4      ///< packed RGB 1:2:1 bitstream,  4bpp, (msb)1R 2G 1B(lsb), a byte contains two pixels, the first pixel in the byte is the one composed by the 4 msb bits
	AV_PIX_FMT_RGB4_BYTE = C.AV_PIX_FMT_RGB4_BYTE ///< packed RGB 1:2:1,  8bpp, (msb)1R 2G 1B(lsb)
	AV_PIX_FMT_NV12      = C.AV_PIX_FMT_NV12      ///< planar YUV 4:2:0, 12bpp, 1 plane for Y and 1 plane for the UV components, which are interleaved (first byte U and the following byte V)
	AV_PIX_FMT_NV21      = C.AV_PIX_FMT_NV21      ///< as above, but U and V bytes are swapped

	AV_PIX_FMT_ARGB = C.AV_PIX_FMT_ARGB ///< packed ARGB 8:8:8:8, 32bpp, ARGBARGB...
	AV_PIX_FMT_RGBA = C.AV_PIX_FMT_RGBA ///< packed RGBA 8:8:8:8, 32bpp, RGBARGBA...
	AV_PIX_FMT_ABGR = C.AV_PIX_FMT_ABGR ///< packed ABGR 8:8:8:8, 32bpp, ABGRABGR...
	AV_PIX_FMT_BGRA = C.AV_PIX_FMT_BGRA ///< packed BGRA 8:8:8:8, 32bpp, BGRABGRA...

	AV_PIX_FMT_GRAY16BE = C.AV_PIX_FMT_GRAY16BE ///<        Y        , 16bpp, big-endian
	AV_PIX_FMT_GRAY16LE = C.AV_PIX_FMT_GRAY16LE ///<        Y        , 16bpp, little-endian
	AV_PIX_FMT_YUV440P  = C.AV_PIX_FMT_YUV440P  ///< planar YUV 4:4:0 (1 Cr & Cb sample per 1x2 Y samples)
	AV_PIX_FMT_YUVJ440P = C.AV_PIX_FMT_YUVJ440P ///< planar YUV 4:4:0 full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV440P and setting color_range
	AV_PIX_FMT_YUVA420P = C.AV_PIX_FMT_YUVA420P ///< planar YUV 4:2:0, 20bpp, (1 Cr & Cb sample per 2x2 Y & A samples)
	AV_PIX_FMT_RGB48BE  = C.AV_PIX_FMT_RGB48BE  ///< packed RGB 16:16:16, 48bpp, 16R, 16G, 16B, the 2-byte value for each R/G/B component is stored as big-endian
	AV_PIX_FMT_RGB48LE  = C.AV_PIX_FMT_RGB48LE  ///< packed RGB 16:16:16, 48bpp, 16R, 16G, 16B, the 2-byte value for each R/G/B component is stored as little-endian

	AV_PIX_FMT_RGB565BE = C.AV_PIX_FMT_RGB565BE ///< packed RGB 5:6:5, 16bpp, (msb)   5R 6G 5B(lsb), big-endian
	AV_PIX_FMT_RGB565LE = C.AV_PIX_FMT_RGB565LE ///< packed RGB 5:6:5, 16bpp, (msb)   5R 6G 5B(lsb), little-endian
	AV_PIX_FMT_RGB555BE = C.AV_PIX_FMT_RGB555BE ///< packed RGB 5:5:5, 16bpp, (msb)1X 5R 5G 5B(lsb), big-endian   , X=unused/undefined
	AV_PIX_FMT_RGB555LE = C.AV_PIX_FMT_RGB555LE ///< packed RGB 5:5:5, 16bpp, (msb)1X 5R 5G 5B(lsb), little-endian, X=unused/undefined

	AV_PIX_FMT_BGR565BE = C.AV_PIX_FMT_BGR565BE ///< packed BGR 5:6:5, 16bpp, (msb)   5B 6G 5R(lsb), big-endian
	AV_PIX_FMT_BGR565LE = C.AV_PIX_FMT_BGR565LE ///< packed BGR 5:6:5, 16bpp, (msb)   5B 6G 5R(lsb), little-endian
	AV_PIX_FMT_BGR555BE = C.AV_PIX_FMT_BGR555BE ///< packed BGR 5:5:5, 16bpp, (msb)1X 5B 5G 5R(lsb), big-endian   , X=unused/undefined
	AV_PIX_FMT_BGR555LE = C.AV_PIX_FMT_BGR555LE ///< packed BGR 5:5:5, 16bpp, (msb)1X 5B 5G 5R(lsb), little-endian, X=unused/undefined

	/**
	 *  Hardware acceleration through VA-API, data[3] contains a
	 *  VASurfaceID.
	 */
	AV_PIX_FMT_VAAPI = C.AV_PIX_FMT_VAAPI

	AV_PIX_FMT_YUV420P16LE = C.AV_PIX_FMT_YUV420P16LE ///< planar YUV 4:2:0, 24bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	AV_PIX_FMT_YUV420P16BE = C.AV_PIX_FMT_YUV420P16BE ///< planar YUV 4:2:0, 24bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	AV_PIX_FMT_YUV422P16LE = C.AV_PIX_FMT_YUV422P16LE ///< planar YUV 4:2:2, 32bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AV_PIX_FMT_YUV422P16BE = C.AV_PIX_FMT_YUV422P16BE ///< planar YUV 4:2:2, 32bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	AV_PIX_FMT_YUV444P16LE = C.AV_PIX_FMT_YUV444P16LE ///< planar YUV 4:4:4, 48bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	AV_PIX_FMT_YUV444P16BE = C.AV_PIX_FMT_YUV444P16BE ///< planar YUV 4:4:4, 48bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	AV_PIX_FMT_DXVA2_VLD   = C.AV_PIX_FMT_DXVA2_VLD   ///< HW decoding through DXVA2, Picture.data[3] contains a LPDIRECT3DSURFACE9 pointer

	AV_PIX_FMT_RGB444LE = C.AV_PIX_FMT_RGB444LE ///< packed RGB 4:4:4, 16bpp, (msb)4X 4R 4G 4B(lsb), little-endian, X=unused/undefined
	AV_PIX_FMT_RGB444BE = C.AV_PIX_FMT_RGB444BE ///< packed RGB 4:4:4, 16bpp, (msb)4X 4R 4G 4B(lsb), big-endian,    X=unused/undefined
	AV_PIX_FMT_BGR444LE = C.AV_PIX_FMT_BGR444LE ///< packed BGR 4:4:4, 16bpp, (msb)4X 4B 4G 4R(lsb), little-endian, X=unused/undefined
	AV_PIX_FMT_BGR444BE = C.AV_PIX_FMT_BGR444BE ///< packed BGR 4:4:4, 16bpp, (msb)4X 4B 4G 4R(lsb), big-endian,    X=unused/undefined
	AV_PIX_FMT_YA8      = C.AV_PIX_FMT_YA8      ///< 8 bits gray, 8 bits alpha

	AV_PIX_FMT_Y400A  = C.AV_PIX_FMT_YA8 ///< alias for AV_PIX_FMT_YA8
	AV_PIX_FMT_GRAY8A = C.AV_PIX_FMT_YA8 ///< alias for AV_PIX_FMT_YA8

	AV_PIX_FMT_BGR48BE = C.AV_PIX_FMT_BGR48BE ///< packed RGB 16:16:16, 48bpp, 16B, 16G, 16R, the 2-byte value for each R/G/B component is stored as big-endian
	AV_PIX_FMT_BGR48LE = C.AV_PIX_FMT_BGR48LE ///< packed RGB 16:16:16, 48bpp, 16B, 16G, 16R, the 2-byte value for each R/G/B component is stored as little-endian

	/**
	 * The following 12 formats have the disadvantage of needing 1 format for each bit depth.
	 * Notice that each 9/10 bits sample is stored in 16 bits with extra padding.
	 * If you want to support multiple bit depths, then using AV_PIX_FMT_YUV420P16* with the bpp stored separately is better.
	 */
	AV_PIX_FMT_YUV420P9BE   = C.AV_PIX_FMT_YUV420P9BE   ///< planar YUV 4:2:0, 13.5bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	AV_PIX_FMT_YUV420P9LE   = C.AV_PIX_FMT_YUV420P9LE   ///< planar YUV 4:2:0, 13.5bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	AV_PIX_FMT_YUV420P10BE  = C.AV_PIX_FMT_YUV420P10BE  ///< planar YUV 4:2:0, 15bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	AV_PIX_FMT_YUV420P10LE  = C.AV_PIX_FMT_YUV420P10LE  ///< planar YUV 4:2:0, 15bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	AV_PIX_FMT_YUV422P10BE  = C.AV_PIX_FMT_YUV422P10BE  ///< planar YUV 4:2:2, 20bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	AV_PIX_FMT_YUV422P10LE  = C.AV_PIX_FMT_YUV422P10LE  ///< planar YUV 4:2:2, 20bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AV_PIX_FMT_YUV444P9BE   = C.AV_PIX_FMT_YUV444P9BE   ///< planar YUV 4:4:4, 27bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	AV_PIX_FMT_YUV444P9LE   = C.AV_PIX_FMT_YUV444P9LE   ///< planar YUV 4:4:4, 27bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	AV_PIX_FMT_YUV444P10BE  = C.AV_PIX_FMT_YUV444P10BE  ///< planar YUV 4:4:4, 30bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	AV_PIX_FMT_YUV444P10LE  = C.AV_PIX_FMT_YUV444P10LE  ///< planar YUV 4:4:4, 30bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	AV_PIX_FMT_YUV422P9BE   = C.AV_PIX_FMT_YUV422P9BE   ///< planar YUV 4:2:2, 18bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	AV_PIX_FMT_YUV422P9LE   = C.AV_PIX_FMT_YUV422P9LE   ///< planar YUV 4:2:2, 18bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AV_PIX_FMT_GBRP         = C.AV_PIX_FMT_GBRP         ///< planar GBR 4:4:4 24bpp
	AV_PIX_FMT_GBR24P       = C.AV_PIX_FMT_GBRP         // alias for #AV_PIX_FMT_GBRP
	AV_PIX_FMT_GBRP9BE      = C.AV_PIX_FMT_GBRP9BE      ///< planar GBR 4:4:4 27bpp, big-endian
	AV_PIX_FMT_GBRP9LE      = C.AV_PIX_FMT_GBRP9LE      ///< planar GBR 4:4:4 27bpp, little-endian
	AV_PIX_FMT_GBRP10BE     = C.AV_PIX_FMT_GBRP10BE     ///< planar GBR 4:4:4 30bpp, big-endian
	AV_PIX_FMT_GBRP10LE     = C.AV_PIX_FMT_GBRP10LE     ///< planar GBR 4:4:4 30bpp, little-endian
	AV_PIX_FMT_GBRP16BE     = C.AV_PIX_FMT_GBRP16BE     ///< planar GBR 4:4:4 48bpp, big-endian
	AV_PIX_FMT_GBRP16LE     = C.AV_PIX_FMT_GBRP16LE     ///< planar GBR 4:4:4 48bpp, little-endian
	AV_PIX_FMT_YUVA422P     = C.AV_PIX_FMT_YUVA422P     ///< planar YUV 4:2:2 24bpp, (1 Cr & Cb sample per 2x1 Y & A samples)
	AV_PIX_FMT_YUVA444P     = C.AV_PIX_FMT_YUVA444P     ///< planar YUV 4:4:4 32bpp, (1 Cr & Cb sample per 1x1 Y & A samples)
	AV_PIX_FMT_YUVA420P9BE  = C.AV_PIX_FMT_YUVA420P9BE  ///< planar YUV 4:2:0 22.5bpp, (1 Cr & Cb sample per 2x2 Y & A samples), big-endian
	AV_PIX_FMT_YUVA420P9LE  = C.AV_PIX_FMT_YUVA420P9LE  ///< planar YUV 4:2:0 22.5bpp, (1 Cr & Cb sample per 2x2 Y & A samples), little-endian
	AV_PIX_FMT_YUVA422P9BE  = C.AV_PIX_FMT_YUVA422P9BE  ///< planar YUV 4:2:2 27bpp, (1 Cr & Cb sample per 2x1 Y & A samples), big-endian
	AV_PIX_FMT_YUVA422P9LE  = C.AV_PIX_FMT_YUVA422P9LE  ///< planar YUV 4:2:2 27bpp, (1 Cr & Cb sample per 2x1 Y & A samples), little-endian
	AV_PIX_FMT_YUVA444P9BE  = C.AV_PIX_FMT_YUVA444P9BE  ///< planar YUV 4:4:4 36bpp, (1 Cr & Cb sample per 1x1 Y & A samples), big-endian
	AV_PIX_FMT_YUVA444P9LE  = C.AV_PIX_FMT_YUVA444P9LE  ///< planar YUV 4:4:4 36bpp, (1 Cr & Cb sample per 1x1 Y & A samples), little-endian
	AV_PIX_FMT_YUVA420P10BE = C.AV_PIX_FMT_YUVA420P10BE ///< planar YUV 4:2:0 25bpp, (1 Cr & Cb sample per 2x2 Y & A samples, big-endian)
	AV_PIX_FMT_YUVA420P10LE = C.AV_PIX_FMT_YUVA420P10LE ///< planar YUV 4:2:0 25bpp, (1 Cr & Cb sample per 2x2 Y & A samples, little-endian)
	AV_PIX_FMT_YUVA422P10BE = C.AV_PIX_FMT_YUVA422P10BE ///< planar YUV 4:2:2 30bpp, (1 Cr & Cb sample per 2x1 Y & A samples, big-endian)
	AV_PIX_FMT_YUVA422P10LE = C.AV_PIX_FMT_YUVA422P10LE ///< planar YUV 4:2:2 30bpp, (1 Cr & Cb sample per 2x1 Y & A samples, little-endian)
	AV_PIX_FMT_YUVA444P10BE = C.AV_PIX_FMT_YUVA444P10BE ///< planar YUV 4:4:4 40bpp, (1 Cr & Cb sample per 1x1 Y & A samples, big-endian)
	AV_PIX_FMT_YUVA444P10LE = C.AV_PIX_FMT_YUVA444P10LE ///< planar YUV 4:4:4 40bpp, (1 Cr & Cb sample per 1x1 Y & A samples, little-endian)
	AV_PIX_FMT_YUVA420P16BE = C.AV_PIX_FMT_YUVA420P16BE ///< planar YUV 4:2:0 40bpp, (1 Cr & Cb sample per 2x2 Y & A samples, big-endian)
	AV_PIX_FMT_YUVA420P16LE = C.AV_PIX_FMT_YUVA420P16LE ///< planar YUV 4:2:0 40bpp, (1 Cr & Cb sample per 2x2 Y & A samples, little-endian)
	AV_PIX_FMT_YUVA422P16BE = C.AV_PIX_FMT_YUVA422P16BE ///< planar YUV 4:2:2 48bpp, (1 Cr & Cb sample per 2x1 Y & A samples, big-endian)
	AV_PIX_FMT_YUVA422P16LE = C.AV_PIX_FMT_YUVA422P16LE ///< planar YUV 4:2:2 48bpp, (1 Cr & Cb sample per 2x1 Y & A samples, little-endian)
	AV_PIX_FMT_YUVA444P16BE = C.AV_PIX_FMT_YUVA444P16BE ///< planar YUV 4:4:4 64bpp, (1 Cr & Cb sample per 1x1 Y & A samples, big-endian)
	AV_PIX_FMT_YUVA444P16LE = C.AV_PIX_FMT_YUVA444P16LE ///< planar YUV 4:4:4 64bpp, (1 Cr & Cb sample per 1x1 Y & A samples, little-endian)

	AV_PIX_FMT_VDPAU = C.AV_PIX_FMT_VDPAU ///< HW acceleration through VDPAU, Picture.data[3] contains a VdpVideoSurface

	AV_PIX_FMT_XYZ12LE = C.AV_PIX_FMT_XYZ12LE ///< packed XYZ 4:4:4, 36 bpp, (msb) 12X, 12Y, 12Z (lsb), the 2-byte value for each X/Y/Z is stored as little-endian, the 4 lower bits are set to 0
	AV_PIX_FMT_XYZ12BE = C.AV_PIX_FMT_XYZ12BE ///< packed XYZ 4:4:4, 36 bpp, (msb) 12X, 12Y, 12Z (lsb), the 2-byte value for each X/Y/Z is stored as big-endian, the 4 lower bits are set to 0
	AV_PIX_FMT_NV16    = C.AV_PIX_FMT_NV16    ///< interleaved chroma YUV 4:2:2, 16bpp, (1 Cr & Cb sample per 2x1 Y samples)
	AV_PIX_FMT_NV20LE  = C.AV_PIX_FMT_NV20LE  ///< interleaved chroma YUV 4:2:2, 20bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AV_PIX_FMT_NV20BE  = C.AV_PIX_FMT_NV20BE  ///< interleaved chroma YUV 4:2:2, 20bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian

	AV_PIX_FMT_RGBA64BE = C.AV_PIX_FMT_RGBA64BE ///< packed RGBA 16:16:16:16, 64bpp, 16R, 16G, 16B, 16A, the 2-byte value for each R/G/B/A component is stored as big-endian
	AV_PIX_FMT_RGBA64LE = C.AV_PIX_FMT_RGBA64LE ///< packed RGBA 16:16:16:16, 64bpp, 16R, 16G, 16B, 16A, the 2-byte value for each R/G/B/A component is stored as little-endian
	AV_PIX_FMT_BGRA64BE = C.AV_PIX_FMT_BGRA64BE ///< packed RGBA 16:16:16:16, 64bpp, 16B, 16G, 16R, 16A, the 2-byte value for each R/G/B/A component is stored as big-endian
	AV_PIX_FMT_BGRA64LE = C.AV_PIX_FMT_BGRA64LE ///< packed RGBA 16:16:16:16, 64bpp, 16B, 16G, 16R, 16A, the 2-byte value for each R/G/B/A component is stored as little-endian

	AV_PIX_FMT_YVYU422 = C.AV_PIX_FMT_YVYU422 ///< packed YUV 4:2:2, 16bpp, Y0 Cr Y1 Cb

	AV_PIX_FMT_YA16BE = C.AV_PIX_FMT_YA16BE ///< 16 bits gray, 16 bits alpha (big-endian)
	AV_PIX_FMT_YA16LE = C.AV_PIX_FMT_YA16LE ///< 16 bits gray, 16 bits alpha (little-endian)

	AV_PIX_FMT_GBRAP     = C.AV_PIX_FMT_GBRAP     ///< planar GBRA 4:4:4:4 32bpp
	AV_PIX_FMT_GBRAP16BE = C.AV_PIX_FMT_GBRAP16BE ///< planar GBRA 4:4:4:4 64bpp, big-endian
	AV_PIX_FMT_GBRAP16LE = C.AV_PIX_FMT_GBRAP16LE ///< planar GBRA 4:4:4:4 64bpp, little-endian
	/**
	 *  HW acceleration through QSV, data[3] contains a pointer to the
	 *  mfxFrameSurface1 structure.
	 */
	AV_PIX_FMT_QSV = C.AV_PIX_FMT_QSV
	/**
	 * HW acceleration though MMAL, data[3] contains a pointer to the
	 * MMAL_BUFFER_HEADER_T structure.
	 */
	AV_PIX_FMT_MMAL = C.AV_PIX_FMT_MMAL

	AV_PIX_FMT_D3D11VA_VLD = C.AV_PIX_FMT_D3D11VA_VLD ///< HW decoding through Direct3D11 via old API, Picture.data[3] contains a ID3D11VideoDecoderOutputView pointer

	/**
	 * HW acceleration through CUDA. data[i] contain CUdeviceptr pointers
	 * exactly as for system memory frames.
	 */
	AV_PIX_FMT_CUDA = C.AV_PIX_FMT_CUDA

	AV_PIX_FMT_0RGB = C.AV_PIX_FMT_0RGB ///< packed RGB 8:8:8, 32bpp, XRGBXRGB...   X=unused/undefined
	AV_PIX_FMT_RGB0 = C.AV_PIX_FMT_RGB0 ///< packed RGB 8:8:8, 32bpp, RGBXRGBX...   X=unused/undefined
	AV_PIX_FMT_0BGR = C.AV_PIX_FMT_0BGR ///< packed BGR 8:8:8, 32bpp, XBGRXBGR...   X=unused/undefined
	AV_PIX_FMT_BGR0 = C.AV_PIX_FMT_BGR0 ///< packed BGR 8:8:8, 32bpp, BGRXBGRX...   X=unused/undefined

	AV_PIX_FMT_YUV420P12BE = C.AV_PIX_FMT_YUV420P12BE ///< planar YUV 4:2:0,18bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	AV_PIX_FMT_YUV420P12LE = C.AV_PIX_FMT_YUV420P12LE ///< planar YUV 4:2:0,18bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	AV_PIX_FMT_YUV420P14BE = C.AV_PIX_FMT_YUV420P14BE ///< planar YUV 4:2:0,21bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	AV_PIX_FMT_YUV420P14LE = C.AV_PIX_FMT_YUV420P14LE ///< planar YUV 4:2:0,21bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	AV_PIX_FMT_YUV422P12BE = C.AV_PIX_FMT_YUV422P12BE ///< planar YUV 4:2:2,24bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	AV_PIX_FMT_YUV422P12LE = C.AV_PIX_FMT_YUV422P12LE ///< planar YUV 4:2:2,24bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AV_PIX_FMT_YUV422P14BE = C.AV_PIX_FMT_YUV422P14BE ///< planar YUV 4:2:2,28bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	AV_PIX_FMT_YUV422P14LE = C.AV_PIX_FMT_YUV422P14LE ///< planar YUV 4:2:2,28bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AV_PIX_FMT_YUV444P12BE = C.AV_PIX_FMT_YUV444P12BE ///< planar YUV 4:4:4,36bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	AV_PIX_FMT_YUV444P12LE = C.AV_PIX_FMT_YUV444P12LE ///< planar YUV 4:4:4,36bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	AV_PIX_FMT_YUV444P14BE = C.AV_PIX_FMT_YUV444P14BE ///< planar YUV 4:4:4,42bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	AV_PIX_FMT_YUV444P14LE = C.AV_PIX_FMT_YUV444P14LE ///< planar YUV 4:4:4,42bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	AV_PIX_FMT_GBRP12BE    = C.AV_PIX_FMT_GBRP12BE    ///< planar GBR 4:4:4 36bpp, big-endian
	AV_PIX_FMT_GBRP12LE    = C.AV_PIX_FMT_GBRP12LE    ///< planar GBR 4:4:4 36bpp, little-endian
	AV_PIX_FMT_GBRP14BE    = C.AV_PIX_FMT_GBRP14BE    ///< planar GBR 4:4:4 42bpp, big-endian
	AV_PIX_FMT_GBRP14LE    = C.AV_PIX_FMT_GBRP14LE    ///< planar GBR 4:4:4 42bpp, little-endian
	AV_PIX_FMT_YUVJ411P    = C.AV_PIX_FMT_YUVJ411P    ///< planar YUV 4:1:1, 12bpp, (1 Cr & Cb sample per 4x1 Y samples) full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV411P and setting color_range

	AV_PIX_FMT_BAYER_BGGR8    = C.AV_PIX_FMT_BAYER_BGGR8    ///< bayer, BGBG..(odd line), GRGR..(even line), 8-bit samples */
	AV_PIX_FMT_BAYER_RGGB8    = C.AV_PIX_FMT_BAYER_RGGB8    ///< bayer, RGRG..(odd line), GBGB..(even line), 8-bit samples */
	AV_PIX_FMT_BAYER_GBRG8    = C.AV_PIX_FMT_BAYER_GBRG8    ///< bayer, GBGB..(odd line), RGRG..(even line), 8-bit samples */
	AV_PIX_FMT_BAYER_GRBG8    = C.AV_PIX_FMT_BAYER_GRBG8    ///< bayer, GRGR..(odd line), BGBG..(even line), 8-bit samples */
	AV_PIX_FMT_BAYER_BGGR16LE = C.AV_PIX_FMT_BAYER_BGGR16LE ///< bayer, BGBG..(odd line), GRGR..(even line), 16-bit samples, little-endian */
	AV_PIX_FMT_BAYER_BGGR16BE = C.AV_PIX_FMT_BAYER_BGGR16BE ///< bayer, BGBG..(odd line), GRGR..(even line), 16-bit samples, big-endian */
	AV_PIX_FMT_BAYER_RGGB16LE = C.AV_PIX_FMT_BAYER_RGGB16LE ///< bayer, RGRG..(odd line), GBGB..(even line), 16-bit samples, little-endian */
	AV_PIX_FMT_BAYER_RGGB16BE = C.AV_PIX_FMT_BAYER_RGGB16BE ///< bayer, RGRG..(odd line), GBGB..(even line), 16-bit samples, big-endian */
	AV_PIX_FMT_BAYER_GBRG16LE = C.AV_PIX_FMT_BAYER_GBRG16LE ///< bayer, GBGB..(odd line), RGRG..(even line), 16-bit samples, little-endian */
	AV_PIX_FMT_BAYER_GBRG16BE = C.AV_PIX_FMT_BAYER_GBRG16BE ///< bayer, GBGB..(odd line), RGRG..(even line), 16-bit samples, big-endian */
	AV_PIX_FMT_BAYER_GRBG16LE = C.AV_PIX_FMT_BAYER_GRBG16LE ///< bayer, GRGR..(odd line), BGBG..(even line), 16-bit samples, little-endian */
	AV_PIX_FMT_BAYER_GRBG16BE = C.AV_PIX_FMT_BAYER_GRBG16BE ///< bayer, GRGR..(odd line), BGBG..(even line), 16-bit samples, big-endian */

	AV_PIX_FMT_YUV440P10LE = C.AV_PIX_FMT_YUV440P10LE ///< planar YUV 4:4:0,20bpp, (1 Cr & Cb sample per 1x2 Y samples), little-endian
	AV_PIX_FMT_YUV440P10BE = C.AV_PIX_FMT_YUV440P10BE ///< planar YUV 4:4:0,20bpp, (1 Cr & Cb sample per 1x2 Y samples), big-endian
	AV_PIX_FMT_YUV440P12LE = C.AV_PIX_FMT_YUV440P12LE ///< planar YUV 4:4:0,24bpp, (1 Cr & Cb sample per 1x2 Y samples), little-endian
	AV_PIX_FMT_YUV440P12BE = C.AV_PIX_FMT_YUV440P12BE ///< planar YUV 4:4:0,24bpp, (1 Cr & Cb sample per 1x2 Y samples), big-endian
	AV_PIX_FMT_AYUV64LE    = C.AV_PIX_FMT_AYUV64LE    ///< packed AYUV 4:4:4,64bpp (1 Cr & Cb sample per 1x1 Y & A samples), little-endian
	AV_PIX_FMT_AYUV64BE    = C.AV_PIX_FMT_AYUV64BE    ///< packed AYUV 4:4:4,64bpp (1 Cr & Cb sample per 1x1 Y & A samples), big-endian

	AV_PIX_FMT_VIDEOTOOLBOX = C.AV_PIX_FMT_VIDEOTOOLBOX ///< hardware decoding through Videotoolbox

	AV_PIX_FMT_P010LE = C.AV_PIX_FMT_P010LE ///< like NV12, with 10bpp per component, data in the high bits, zeros in the low bits, little-endian
	AV_PIX_FMT_P010BE = C.AV_PIX_FMT_P010BE ///< like NV12, with 10bpp per component, data in the high bits, zeros in the low bits, big-endian

	AV_PIX_FMT_GBRAP12BE = C.AV_PIX_FMT_GBRAP12BE ///< planar GBR 4:4:4:4 48bpp, big-endian
	AV_PIX_FMT_GBRAP12LE = C.AV_PIX_FMT_GBRAP12LE ///< planar GBR 4:4:4:4 48bpp, little-endian

	AV_PIX_FMT_GBRAP10BE = C.AV_PIX_FMT_GBRAP10BE ///< planar GBR 4:4:4:4 40bpp, big-endian
	AV_PIX_FMT_GBRAP10LE = C.AV_PIX_FMT_GBRAP10LE ///< planar GBR 4:4:4:4 40bpp, little-endian

	AV_PIX_FMT_MEDIACODEC = C.AV_PIX_FMT_MEDIACODEC ///< hardware decoding through MediaCodec

	AV_PIX_FMT_GRAY12BE = C.AV_PIX_FMT_GRAY12BE ///<        Y        , 12bpp, big-endian
	AV_PIX_FMT_GRAY12LE = C.AV_PIX_FMT_GRAY12LE ///<        Y        , 12bpp, little-endian
	AV_PIX_FMT_GRAY10BE = C.AV_PIX_FMT_GRAY10BE ///<        Y        , 10bpp, big-endian
	AV_PIX_FMT_GRAY10LE = C.AV_PIX_FMT_GRAY10LE ///<        Y        , 10bpp, little-endian

	AV_PIX_FMT_P016LE = C.AV_PIX_FMT_P016LE ///< like NV12, with 16bpp per component, little-endian
	AV_PIX_FMT_P016BE = C.AV_PIX_FMT_P016BE ///< like NV12, with 16bpp per component, big-endian

	/**
	 * Hardware surfaces for Direct3D11.
	 *
	 * This is preferred over the legacy AV_PIX_FMT_D3D11VA_VLD. The new D3D11
	 * hwaccel API and filtering support AV_PIX_FMT_D3D11 only.
	 *
	 * data[0] contains a ID3D11Texture2D pointer, and data[1] contains the
	 * texture array index of the frame as intptr_t if the ID3D11Texture2D is
	 * an array texture (or always 0 if it's a normal texture).
	 */
	AV_PIX_FMT_D3D11 = C.AV_PIX_FMT_D3D11

	AV_PIX_FMT_GRAY9BE = C.AV_PIX_FMT_GRAY9BE ///<        Y        , 9bpp, big-endian
	AV_PIX_FMT_GRAY9LE = C.AV_PIX_FMT_GRAY9LE ///<        Y        , 9bpp, little-endian

	AV_PIX_FMT_GBRPF32BE  = C.AV_PIX_FMT_GBRPF32BE  ///< IEEE-754 single precision planar GBR 4:4:4,     96bpp, big-endian
	AV_PIX_FMT_GBRPF32LE  = C.AV_PIX_FMT_GBRPF32LE  ///< IEEE-754 single precision planar GBR 4:4:4,     96bpp, little-endian
	AV_PIX_FMT_GBRAPF32BE = C.AV_PIX_FMT_GBRAPF32BE ///< IEEE-754 single precision planar GBRA 4:4:4:4, 128bpp, big-endian
	AV_PIX_FMT_GBRAPF32LE = C.AV_PIX_FMT_GBRAPF32LE ///< IEEE-754 single precision planar GBRA 4:4:4:4, 128bpp, little-endian

	/**
	 * DRM-managed buffers exposed through PRIME buffer sharing.
	 *
	 * data[0] points to an AVDRMFrameDescriptor.
	 */
	AV_PIX_FMT_DRM_PRIME = C.AV_PIX_FMT_DRM_PRIME
	/**
	 * Hardware surfaces for OpenCL.
	 *
	 * data[i] contain 2D image objects (typed in C as cl_mem, used
	 * in OpenCL as image2d_t) for each plane of the surface.
	 */
	AV_PIX_FMT_OPENCL = C.AV_PIX_FMT_OPENCL

	AV_PIX_FMT_GRAY14BE = C.AV_PIX_FMT_GRAY14BE ///<        Y        , 14bpp, big-endian
	AV_PIX_FMT_GRAY14LE = C.AV_PIX_FMT_GRAY14LE ///<        Y        , 14bpp, little-endian

	AV_PIX_FMT_GRAYF32BE = C.AV_PIX_FMT_GRAYF32BE ///< IEEE-754 single precision Y, 32bpp, big-endian
	AV_PIX_FMT_GRAYF32LE = C.AV_PIX_FMT_GRAYF32LE ///< IEEE-754 single precision Y, 32bpp, little-endian

	// Original constants from goav
	AV_PIX_FMT_YUV        = C.AV_PIX_FMT_YUV420P
	AV_PIX_FMT_YUV420P9   = C.AV_PIX_FMT_YUV420P9
	AV_PIX_FMT_YUV422P9   = C.AV_PIX_FMT_YUV422P9
	AV_PIX_FMT_YUV444P9   = C.AV_PIX_FMT_YUV444P9
	AV_PIX_FMT_YUV420P10  = C.AV_PIX_FMT_YUV420P10
	AV_PIX_FMT_YUV422P10  = C.AV_PIX_FMT_YUV422P10
	AV_PIX_FMT_YUV440P10  = C.AV_PIX_FMT_YUV440P10
	AV_PIX_FMT_YUV444P10  = C.AV_PIX_FMT_YUV444P10
	AV_PIX_FMT_YUV420P12  = C.AV_PIX_FMT_YUV420P12
	AV_PIX_FMT_YUV422P12  = C.AV_PIX_FMT_YUV422P12
	AV_PIX_FMT_YUV440P12  = C.AV_PIX_FMT_YUV440P12
	AV_PIX_FMT_YUV444P12  = C.AV_PIX_FMT_YUV444P12
	AV_PIX_FMT_YUV420P14  = C.AV_PIX_FMT_YUV420P14
	AV_PIX_FMT_YUV422P14  = C.AV_PIX_FMT_YUV422P14
	AV_PIX_FMT_YUV444P14  = C.AV_PIX_FMT_YUV444P14
	AV_PIX_FMT_YUV420P16  = C.AV_PIX_FMT_YUV420P16
	AV_PIX_FMT_YUV422P16  = C.AV_PIX_FMT_YUV422P16
	AV_PIX_FMT_YUV444P16  = C.AV_PIX_FMT_YUV444P16
	AV_PIX_FMT_YUVA420P9  = C.AV_PIX_FMT_YUVA420P9
	AV_PIX_FMT_YUVA422P9  = C.AV_PIX_FMT_YUVA422P9
	AV_PIX_FMT_YUVA444P9  = C.AV_PIX_FMT_YUVA444P9
	AV_PIX_FMT_YUVA420P10 = C.AV_PIX_FMT_YUVA420P10
	AV_PIX_FMT_YUVA422P10 = C.AV_PIX_FMT_YUVA422P10
	AV_PIX_FMT_YUVA444P10 = C.AV_PIX_FMT_YUVA444P10
	AV_PIX_FMT_YUVA420P16 = C.AV_PIX_FMT_YUVA420P16
	AV_PIX_FMT_YUVA422P16 = C.AV_PIX_FMT_YUVA422P16
	AV_PIX_FMT_YUVA444P16 = C.AV_PIX_FMT_YUVA444P16

	SWS_FAST_BILINEAR        = C.SWS_FAST_BILINEAR
	SWS_BILINEAR             = C.SWS_BILINEAR
	SWS_BICUBIC              = C.SWS_BICUBIC
	SWS_X                    = C.SWS_X
	SWS_POINT                = C.SWS_POINT
	SWS_AREA                 = C.SWS_AREA
	SWS_BICUBLIN             = C.SWS_BICUBLIN
	SWS_GAUSS                = C.SWS_GAUSS
	SWS_SINC                 = C.SWS_SINC
	SWS_LANCZOS              = C.SWS_LANCZOS
	SWS_SPLINE               = C.SWS_SPLINE
	SWS_SRC_V_CHR_DROP_MASK  = C.SWS_SRC_V_CHR_DROP_MASK
	SWS_SRC_V_CHR_DROP_SHIFT = C.SWS_SRC_V_CHR_DROP_SHIFT
	SWS_PARAM_DEFAULT        = C.SWS_PARAM_DEFAULT
	SWS_PRINT_INFO           = C.SWS_PRINT_INFO
	SWS_FULL_CHR_H_INT       = C.SWS_FULL_CHR_H_INT
	SWS_FULL_CHR_H_INP       = C.SWS_FULL_CHR_H_INP
	SWS_DIRECT_BGR           = C.SWS_DIRECT_BGR
	SWS_ACCURATE_RND         = C.SWS_ACCURATE_RND
	SWS_BITEXACT             = C.SWS_BITEXACT
	SWS_ERROR_DIFFUSION      = C.SWS_ERROR_DIFFUSION
	SWS_MAX_REDUCE_CUTOFF    = C.SWS_MAX_REDUCE_CUTOFF
	SWS_CS_ITU709            = C.SWS_CS_ITU709
	SWS_CS_FCC               = C.SWS_CS_FCC
	SWS_CS_ITU601            = C.SWS_CS_ITU601
	SWS_CS_ITU624            = C.SWS_CS_ITU624
	SWS_CS_SMPTE170M         = C.SWS_CS_SMPTE170M
	SWS_CS_SMPTE240M         = C.SWS_CS_SMPTE240M
	SWS_CS_DEFAULT           = C.SWS_CS_DEFAULT
	SWS_CS_BT2020            = C.SWS_CS_BT2020
)

func (pf PixelFormat) String() string {
	switch int(pf) {
	case AV_PIX_FMT_YUV420P9:
		return "YUV420P9"

	case AV_PIX_FMT_YUV422P9:
		return "YUV422P9"

	case AV_PIX_FMT_YUV444P9:
		return "YUV444P9"

	case AV_PIX_FMT_YUV420P10:
		return "YUV420P10"

	case AV_PIX_FMT_YUV422P10:
		return "YUV422P10"

	case AV_PIX_FMT_YUV440P10:
		return "YUV440P10"

	case AV_PIX_FMT_YUV444P10:
		return "YUV444P10"

	case AV_PIX_FMT_YUV420P12:
		return "YUV420P12"

	case AV_PIX_FMT_YUV422P12:
		return "YUV422P12"

	case AV_PIX_FMT_YUV440P12:
		return "YUV440P12"

	case AV_PIX_FMT_YUV444P12:
		return "YUV444P12"

	case AV_PIX_FMT_YUV420P14:
		return "YUV420P14"

	case AV_PIX_FMT_YUV422P14:
		return "YUV422P14"

	case AV_PIX_FMT_YUV444P14:
		return "YUV444P14"

	case AV_PIX_FMT_YUV420P16:
		return "YUV420P16"

	case AV_PIX_FMT_YUV422P16:
		return "YUV422P16"

	case AV_PIX_FMT_YUV444P16:
		return "YUV444P16"

	case AV_PIX_FMT_YUVA420P9:
		return "YUVA420P9"

	case AV_PIX_FMT_YUVA422P9:
		return "YUVA422P9"

	case AV_PIX_FMT_YUVA444P9:
		return "YUVA444P9"

	case AV_PIX_FMT_YUVA420P10:
		return "YUVA420P10"

	case AV_PIX_FMT_YUVA422P10:
		return "YUVA422P10"

	case AV_PIX_FMT_YUVA444P10:
		return "YUVA444P10"

	case AV_PIX_FMT_YUVA420P16:
		return "YUVA420P16"

	case AV_PIX_FMT_YUVA422P16:
		return "YUVA422P16"

	case AV_PIX_FMT_YUVA444P16:
		return "YUVA444P16"

	case AV_PIX_FMT_RGB24:
		return "RGB24"

	case AV_PIX_FMT_RGBA:
		return "RGBA"
	}

	return "{UNKNOWN}"
}
